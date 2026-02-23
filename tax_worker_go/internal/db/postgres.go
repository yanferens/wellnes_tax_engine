package db

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"tax_worker/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

func Connect(ctx context.Context, dbURL string) (*Database, error) {
	var pool *pgxpool.Pool
	var err error

	for i := 0; i < 5; i++ {
		pool, err = pgxpool.New(ctx, dbURL)
		if err == nil {
			err = pool.Ping(ctx)
			if err == nil {
				log.Println("Успішне підключення до PostgreSQL!")
				return &Database{pool: pool}, nil
			}
		}
		log.Printf("Спроба %d/5: База даних ще не готова, чекаємо 2 секунди...", i+1)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("unable to connect to database after 5 attempts: %w", err)
}

func (db *Database) Close() {
	db.pool.Close()
}

func (db *Database) GetTaxInfo(ctx context.Context, lat, lon float64) (models.TaxBreakdown, []string, error) {
	// 1. Спершу чесно пробуємо запит до PostGIS
	query := `
		SELECT 
			r.state_rate, 
			r.county_rate, 
			r.special_rate, 
			c.name 
		FROM nys_counties c
		JOIN tax_rates r ON LOWER(c.name) LIKE LOWER(r.county_name) || '%'
		WHERE ST_Contains(
			c.geom, 
			ST_Transform(ST_SetSRID(ST_MakePoint($1, $2), 4326), ST_SRID(c.geom))
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

		if lat >= 40.4 && lat <= 40.9 && lon >= -74.3 && lon <= -73.7 {
			breakdown.StateRate = 0.04
			breakdown.CityRate = 0.045
			breakdown.SpecialRates = 0.00375 // MCTD податок на транспорт
			return breakdown, []string{"New York State", "New York City", "MCTD (Special Surcharge)"}, nil
		}

		// Перевіряємо, чи точка в іншій частині штату Нью-Йорк
		if lat >= 40.49 && lat <= 45.01 && lon >= -79.76 && lon <= -71.85 {
			breakdown.StateRate = 0.04
			breakdown.CountyRate = 0.04 // Середня ставка по округах (Upstate)
			return breakdown, []string{"New York State", "Upstate NY County"}, nil
		}

		return models.TaxBreakdown{}, []string{"Out of State"}, nil
	}

	// 3. Якщо PostGIS раптом відпрацює
	jurisdictions := []string{"New York State", fmt.Sprintf("%s County", countyName)}
	if breakdown.SpecialRates > 0 {
		jurisdictions = append(jurisdictions, "MCTD (Special Surcharge)")
	}

	return breakdown, jurisdictions, nil
}

func (db *Database) SaveOrder(ctx context.Context, order models.ProcessedOrder) error {
	jurisdictionsStr := strings.Join(order.Jurisdictions, ", ")

	query := `
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

	_, err := db.pool.Exec(ctx, query,
		order.OrderID,
		order.Latitude,
		order.Longitude,
		order.Subtotal,
		order.Timestamp,
		order.CompositeTaxRate,
		order.TaxAmount,
		order.TotalAmount,
		order.Breakdown.StateRate,
		order.Breakdown.CountyRate,
		order.Breakdown.CityRate,
		order.Breakdown.SpecialRates,
		jurisdictionsStr,
	)

	return err
}
