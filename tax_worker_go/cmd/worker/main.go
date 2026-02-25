package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"tax_worker/internal/config"
	"tax_worker/internal/db"
	"tax_worker/internal/models"
	"tax_worker/internal/queue"
	"tax_worker/internal/tax"
	"time"
)

func main() {
	cfg := config.Load()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	q := queue.NewRedisQueue(cfg.RedisURL)
	ordersChan, err := q.Consume(ctx)
	if err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	const numWorkers = 100
	var wg sync.WaitGroup

	var processedCount int64
	var batchStartTime time.Time
	var isProcessing int32
	var mu sync.Mutex

	totalToProcess := int64(11222)

	resultsChan := make(chan models.ProcessedOrder, 2000)

	log.Printf("🚀 Tax Worker started with %d workers. Ready for parallel processing...", numWorkers)

	go func() {
		var batch []models.ProcessedOrder
		batchSize := 1000
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case res := <-resultsChan:
				batch = append(batch, res)
				if len(batch) >= batchSize {

					if err := database.SaveOrdersBatch(context.Background(), batch); err != nil {
						log.Printf("Batch save error: %v", err)
					}
					batch = batch[:0]
				}
			case <-ticker.C:
				if len(batch) > 0 {
					if err := database.SaveOrdersBatch(context.Background(), batch); err != nil {
						log.Printf("Batch save error: %v", err)
					}
					batch = batch[:0]
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for order := range ordersChan {

				if atomic.CompareAndSwapInt32(&isProcessing, 0, 1) {
					mu.Lock()
					batchStartTime = time.Now()
					atomic.StoreInt64(&processedCount, 0)
					mu.Unlock()
					log.Println("⏳ Початок обробки нової пачки файлів...")
				}

				breakdown, jurisdictions, err := database.GetTaxInfo(ctx, order.Latitude, order.Longitude)
				if err != nil {
					log.Printf("[Worker %d] Error: %v", workerID, err)
					continue
				}

				result := tax.Calculate(order.Subtotal, breakdown)
				result.OrderID = order.OrderID
				result.Latitude = order.Latitude
				result.Longitude = order.Longitude
				result.Subtotal = order.Subtotal
				result.Timestamp = order.Timestamp
				result.Jurisdictions = jurisdictions

				// 3. Замість збереження по одному, відправляємо в канал
				resultsChan <- result

				newCount := atomic.AddInt64(&processedCount, 1)

				if newCount%1000 == 0 || newCount == totalToProcess {
					mu.Lock()
					duration := time.Since(batchStartTime)
					mu.Unlock()
					log.Printf("📊 Прогрес: [%d/%d] | Час: %v | RPS: %.2f",
						newCount,
						totalToProcess,
						duration,
						float64(newCount)/duration.Seconds(),
					)
				}

				if newCount == totalToProcess {
					mu.Lock()
					finalDuration := time.Since(batchStartTime)
					mu.Unlock()
					log.Printf("✨ ФІНАЛ: Оброблено %d рядків за %v!", newCount, finalDuration)
					atomic.StoreInt32(&isProcessing, 0)
				}
			}
		}(i)
	}

	<-sigChan
	log.Println("🛑 Зупинка воркера...")
	cancel()
	wg.Wait()
}
