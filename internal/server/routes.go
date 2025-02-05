package server

import (
	"github.com/labstack/echo/v4"
	"github.com/lakelimbo/pulsar/internal/server/handlers"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
)

func SetupRoutes(e *echo.Echo, db *bun.DB, redis *redis.Client) {
	// Initialize handlers
	healthHandler := handlers.NewHealthHandler(db, redis)

	// Endpoints
	e.GET("/health", healthHandler.HealthCheck)
}
