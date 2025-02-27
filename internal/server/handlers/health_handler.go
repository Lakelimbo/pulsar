package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
)

type HealthHandler struct {
	db    *bun.DB
	redis *redis.Client
}

func NewHealthHandler(db *bun.DB, redis *redis.Client) *HealthHandler {
	return &HealthHandler{
		db:    db,
		redis: redis,
	}
}

// HealthCheck godoc
//
//	@Summary		Server health check
//	@Description	Checks the health of the server
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		503	{object}	map[string]string
//	@Router			/health [get]
func (h *HealthHandler) HealthCheck(c echo.Context) error {
	ctx := c.Request().Context()

	// Check Postgres connection
	if err := h.db.Ping(); err != nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"status":  "unhealthy",
			"message": "PostgreSQL connection failed",
		})
	}

	// Check Redis connection
	if err := h.redis.Ping(ctx).Err(); err != nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"status":  "unhealthy",
			"message": "Redis connection failed",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status":  "healthy",
		"message": "All systems operational",
	})
}
