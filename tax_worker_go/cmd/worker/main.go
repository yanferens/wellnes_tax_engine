package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"tax_worker/internal/config"
	"tax_worker/internal/db"
	"tax_worker/internal/models"
	"tax_worker/internal/queue"
	"tax_worker/internal/tax"
)

const (
	NumWorkers    = 100
	BatchSize     = 1000
	FlushInterval = 1 * time.Second
)

func main() {
	cfg := config.Load()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("❌ Database connection error: %v", err)
	}
	defer database.Close()

	q := queue.NewRedisQueue(cfg.RedisURL)
	ordersChan, err := q.Consume(ctx)
	if err != nil {
		log.Fatalf("❌ Queue connection error: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	var processedCount int64

	var startTime time.Time
	var startOnce sync.Once

	resultsChan := make(chan models.ProcessedOrder, BatchSize*2)

	log.Printf("🚀 Tax Worker started. Workers: %d, Batch Size: %d", NumWorkers, BatchSize)

	go func() {
		batch := make([]models.ProcessedOrder, 0, BatchSize)
		ticker := time.NewTicker(FlushInterval)
		defer ticker.Stop()

		flush := func() {
			if len(batch) > 0 {
				if err := database.SaveOrdersBatch(context.Background(), batch); err != nil {
					log.Printf("❌ Batch save error: %v", err)
				}
				batch = batch[:0]
			}
		}

		for {
			select {
			case res := <-resultsChan:
				batch = append(batch, res)
				if len(batch) >= BatchSize {
					flush()
				}
			case <-ticker.C:
				flush()
			case <-ctx.Done():
				flush()
				return
			}
		}
	}()

	for i := 0; i < NumWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for order := range ordersChan {

				startOnce.Do(func() {
					startTime = time.Now()
					log.Println("⏳ First order received! Starting timer...")
				})

				breakdown, jurisdictions, err := database.GetTaxInfo(ctx, order.Latitude, order.Longitude)
				if err != nil {
					log.Printf("[Worker %d] Geo-data processing error for OrderID %d: %v", workerID, order.OrderID, err)
					continue
				}

				result := tax.Calculate(order.Subtotal, breakdown)
				result.OrderID = order.OrderID
				result.Latitude = order.Latitude
				result.Longitude = order.Longitude
				result.Subtotal = order.Subtotal
				result.Timestamp = order.Timestamp
				result.Jurisdictions = jurisdictions

				resultsChan <- result

				newCount := atomic.AddInt64(&processedCount, 1)

				if newCount%1000 == 0 {
					duration := time.Since(startTime)
					rps := float64(newCount) / duration.Seconds()
					log.Printf("📊 Progress: %d orders processed | Time: %v | RPS: %.2f", newCount, duration, rps)
				}
			}
		}(i)
	}

	<-sigChan
	log.Println("🛑 Stop signal received. Shutting down...")

	cancel()
	wg.Wait()

	if !startTime.IsZero() {
		finalDuration := time.Since(startTime)
		finalCount := atomic.LoadInt64(&processedCount)
		log.Printf("✨ FINAL: Successfully processed %d rows in %v!", finalCount, finalDuration)
	}
}
