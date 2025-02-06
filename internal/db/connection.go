package db

import (
	"database/sql"
	"fmt"

	"github.com/lakelimbo/pulsar/internal/config"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func NewDBConnection(cfg *config.DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func RunMigrations(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "internal/db/migrations"); err != nil {
		return err
	}

	return nil
}
