#!/bin/bash
set -e

echo "⏳ Starting Database Initialization..."


echo "⏳ Auto-transforming Shapefiles (UTM Meters -> WGS84 Degrees)..."
shp2pgsql -I -s 26918:4326 /geo_data/Counties.shp public.nys_counties | psql -U "$POSTGRES_USER" -d "$POSTGRES_DB"
shp2pgsql -I -s 26918:4326 /geo_data/Cities.shp public.nys_cities | psql -U "$POSTGRES_USER" -d "$POSTGRES_DB"

echo "⏳ Setting up schema and admin user..."

psql -v ON_ERROR_STOP=1 -U "$POSTGRES_USER" -d "$POSTGRES_DB" <<-'EOSQL'
    CREATE EXTENSION IF NOT EXISTS pgcrypto;

    CREATE TABLE IF NOT EXISTS tax_rates (
        county_name VARCHAR(100) PRIMARY KEY,
        state_rate NUMERIC(5,4),
        county_rate NUMERIC(5,4),
        special_rate NUMERIC(5,4)
    );

    INSERT INTO tax_rates (county_name, state_rate, county_rate, special_rate) VALUES
    ('Kings', 0.0400, 0.0450, 0.00375),
    ('New York', 0.0400, 0.0450, 0.00375),
    ('Queens', 0.0400, 0.0450, 0.00375),
    ('Bronx', 0.0400, 0.0450, 0.00375),
    ('Richmond', 0.0400, 0.0450, 0.00375),
    ('Albany', 0.0400, 0.0400, 0.0000),
    ('Westchester', 0.0400, 0.0400, 0.00375)
    ON CONFLICT (county_name) DO NOTHING;

    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) UNIQUE NOT NULL,
        hashed_password VARCHAR(255) NOT NULL
    );

    INSERT INTO users (id, username, hashed_password) VALUES
    (1, 'admin@gmail.com', crypt('admin', gen_salt('bf', 12)))
    ON CONFLICT (id) DO NOTHING;
EOSQL

echo "✅ Database ready and coordinates converted!"