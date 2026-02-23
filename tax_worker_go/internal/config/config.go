package config

import "os"

type Config struct {
	DatabaseURL string
	RedisURL    string
}

func Load() *Config {
	return &Config{
		// os.Getenv бере значення, які Docker підставив із .env
		DatabaseURL: os.Getenv("DATABASE_URL"),
		RedisURL:    os.Getenv("REDIS_URL"),
	}
}
