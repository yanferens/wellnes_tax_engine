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

	const numWorkers = 20
	var wg sync.WaitGroup
	var processedCount int64
	var batchStartTime time.Time
	var once sync.Once
	totalToProcess := int64(11222)

	log.Printf("🚀 Tax Worker started with %d workers. Ready for parallel processing...", numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for order := range ordersChan {

				once.Do(func() {
					batchStartTime = time.Now()
					log.Println("⏳ Початок паралельної обробки файлу...")
				})

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

				if err := database.SaveOrder(ctx, result); err != nil {
					log.Printf("[Worker %d] Save error: %v", workerID, err)
					continue
				}

				newCount := atomic.AddInt64(&processedCount, 1)

				if newCount%500 == 0 || newCount == totalToProcess {
					duration := time.Since(batchStartTime)
					log.Printf("📊 Прогрес: [%d/%d] | Затрачено часу: %v | RPS: %.2f",
						newCount,
						totalToProcess,
						duration,
						float64(newCount)/duration.Seconds(),
					)
				}

				if newCount == totalToProcess {
					log.Printf("✨ ФІНАЛ: Оброблено %d рядків за %v!", newCount, time.Since(batchStartTime))
				}
			}
		}(i)
	}

	<-sigChan
	log.Println("🛑 Зупинка воркера...")
	cancel()
	wg.Wait()
}
