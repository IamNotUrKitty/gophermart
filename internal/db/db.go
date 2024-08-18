package db

import (
	"context"
	"embed"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*
var fs embed.FS

func NewConnectionPool(cs string) (*pgxpool.Pool, error) {
	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, cs)
	if err != nil {
		return nil, err
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return nil, err
		}
	}

	pool, err := pgxpool.New(context.Background(), cs)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
