package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DB    DBConfig
	Redis RedisConfig
	OTel  OTelConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type OTelConfig struct {
	CollectorEndpoint string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	return &Config{
		Port: os.Getenv("PORT"),
		DB: DBConfig{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DB:       os.Getenv("POSTGRES_DB"),
		},
		Redis: RedisConfig{
			Addr:     os.Getenv("REDIS_URL"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       redisDB,
		},
		OTel: OTelConfig{
			CollectorEndpoint: os.Getenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT"),
		},
	}, nil
}
