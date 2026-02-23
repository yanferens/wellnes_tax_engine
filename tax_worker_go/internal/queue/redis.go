package queue

import (
	"context"
	"encoding/json"
	"log"
	"tax_worker/internal/models"

	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	client *redis.Client
}

func NewRedisQueue(redisURL string) *RedisQueue {
	return &RedisQueue{
		client: redis.NewClient(&redis.Options{
			Addr: redisURL,
		}),
	}
}

func (q *RedisQueue) Consume(ctx context.Context) (<-chan models.Order, error) {
	ordersChan := make(chan models.Order)

	if err := q.client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	log.Println("Успішне підключення до Redis!")

	go func() {
		defer close(ordersChan)
		for {
			select {
			case <-ctx.Done():
				return
			default:

				res, err := q.client.BLPop(ctx, 0, "orders_queue").Result()
				if err != nil {
					log.Printf("Помилка читання з Redis: %v", err)
					continue
				}

				if len(res) < 2 {
					continue
				}

				var order models.Order
				if err := json.Unmarshal([]byte(res[1]), &order); err != nil {
					log.Printf("Помилка парсингу JSON з Redis: %v. Дані: %s", err, res[1])
					continue
				}

				ordersChan <- order
			}
		}
	}()

	return ordersChan, nil
}
