package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lakelimbo/pulsar/internal/config"
	"github.com/lakelimbo/pulsar/internal/db"
	"github.com/lakelimbo/pulsar/internal/redis"
	"github.com/lakelimbo/pulsar/internal/server"
	"github.com/lakelimbo/pulsar/internal/tracing"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	tp, err := tracing.InitTracer(&cfg.OTel, "pulsar")
	if err != nil {
		log.Fatalf("Failed to initialize tracer: %v", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	dbConn, err := db.NewDBConnection(&cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	defer dbConn.Close()

	redisClient := redis.NewRedisClient(&cfg.Redis)
	defer redisClient.Close()
	if err := redis.PingRedis(context.Background(), redisClient); err != nil {
		log.Fatalf("Failed to ping Redis: %v", err)
	}

	e := server.NewServer(dbConn, redisClient)

	go func() {
		if err := e.Start(":" + cfg.Port); err != nil {
			e.Logger.Info("Shutting down the server...")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
