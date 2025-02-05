package server

import (
	"github.com/labstack/echo/v4"
	"github.com/lakelimbo/pulsar/internal/server/handlers"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
)

//	@title			Pulsar API
//	@version		0.0.1
//	@description	A cool Go API
//	@host			localhost:8080
//	@BasePath		/api/v1
func SetupRoutes(echo *echo.Echo, db *bun.DB, redis *redis.Client) {
	// Initialize handlers
	healthHandler := handlers.NewHealthHandler(db, redis)

	// Endpoints
	e := echo.Group("/api/v1")

	e.GET("/health", healthHandler.HealthCheck)

	// Scalar
	e.GET("/swagger.yaml", handlers.ServeSwagger)
	e.GET("/reference", handlers.InitScalar)
}
