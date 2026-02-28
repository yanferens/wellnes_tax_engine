package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"tax_worker/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const upsertOrderQuery = `
	INSERT INTO orders (
		id, latitude, longitude, subtotal, timestamp,
		composite_tax_rate, tax_amount, total_amount,
		state_rate, county_rate, city_rate, special_rates,
		jurisdictions
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	ON CONFLICT (id) DO UPDATE SET
		composite_tax_rate = EXCLUDED.composite_tax_rate,
		tax_amount = EXCLUDED.tax_amount,
		total_amount = EXCLUDED.total_amount,
		jurisdictions = EXCLUDED.jurisdictions;
`

type Database struct {
	pool *pgxpool.Pool
}

func Connect(ctx context.Context, dbURL string) (*Database, error) {
	var pool *pgxpool.Pool
	var err error

	for i := 0; i < 5; i++ {
		pool, err = pgxpool.New(ctx, dbURL)
		if err == nil {
			if err = pool.Ping(ctx); err == nil {
				log.Println("✅ Successfully connected to PostgreSQL!")
				return &Database{pool: pool}, nil
			}
		}
		log.Printf("⏳ Attempt %d/5: Database not ready, waiting 2 seconds...", i+1)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("unable to connect to database after 5 attempts: %w", err)
}

func (db *Database) Close() {
	if db.pool != nil {
		db.pool.Close()
	}
}

func (db *Database) GetTaxInfo(ctx context.Context, lat, lon float64) (models.TaxBreakdown, []string, error) {
	query := `
       SELECT 
          r.state_rate, 
          r.county_rate, 
          r.special_rate, 
          c.name 
       FROM nys_counties c
       JOIN tax_rates r ON LOWER(c.name) = LOWER(r.county_name)
       WHERE ST_Contains(
          c.geom, 
          ST_SetSRID(ST_MakePoint($1, $2), 4326)
       ) LIMIT 1;
    `

	var breakdown models.TaxBreakdown
	var countyName string

	err := db.pool.QueryRow(ctx, query, lon, lat).Scan(
		&breakdown.StateRate,
		&breakdown.CountyRate,
		&breakdown.SpecialRates,
		&countyName,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.TaxBreakdown{}, []string{"Out of State"}, nil
		}
		return models.TaxBreakdown{}, nil, fmt.Errorf("database query error: %w", err)
	}

	jurisdictions := []string{"New York State", fmt.Sprintf("%s County", countyName)}
	if breakdown.SpecialRates > 0 {
		jurisdictions = append(jurisdictions, "MCTD (Special Surcharge)")
	}

	return breakdown, jurisdictions, nil
}

func (db *Database) SaveOrder(ctx context.Context, order models.ProcessedOrder) error {
	jurisdictionsStr := strings.Join(order.Jurisdictions, ", ")

	_, err := db.pool.Exec(ctx, upsertOrderQuery,
		order.OrderID, order.Latitude, order.Longitude, order.Subtotal, order.Timestamp,
		order.CompositeTaxRate, order.TaxAmount, order.TotalAmount,
		order.Breakdown.StateRate, order.Breakdown.CountyRate, order.Breakdown.CityRate, order.Breakdown.SpecialRates,
		jurisdictionsStr,
	)
	return err
}

func (db *Database) SaveOrdersBatch(ctx context.Context, orders []models.ProcessedOrder) error {
	if len(orders) == 0 {
		return nil
	}

	batch := &pgx.Batch{}
	for _, order := range orders {
		jurisdictionsStr := strings.Join(order.Jurisdictions, ", ")
		batch.Queue(upsertOrderQuery,
			order.OrderID, order.Latitude, order.Longitude, order.Subtotal, order.Timestamp,
			order.CompositeTaxRate, order.TaxAmount, order.TotalAmount,
			order.Breakdown.StateRate, order.Breakdown.CountyRate, order.Breakdown.CityRate, order.Breakdown.SpecialRates,
			jurisdictionsStr,
		)
	}

	br := db.pool.SendBatch(ctx, batch)
	defer br.Close()

	for i := 0; i < len(orders); i++ {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("batch write error at index %d: %w", i, err)
		}
	}

	return nil
}
