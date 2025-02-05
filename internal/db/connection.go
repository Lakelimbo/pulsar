package db

import (
	"database/sql"
	"fmt"

	"github.com/lakelimbo/pulsar/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewDBConnection(cfg *config.DBConfig) (*bun.DB, error) {
	pg := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)

	sqldb, err := sql.Open("postgres", pg)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
