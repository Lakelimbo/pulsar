package server

import (
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
)

func NewServer(db *bun.DB, redis *redis.Client) *echo.Echo {
	e := echo.New()

	SetupRoutes(e, db, redis)
	return e
}
