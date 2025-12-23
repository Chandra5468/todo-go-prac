package postgresql

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	var ctx context.Context = context.Background()

	// var config *pgxpool.Pool
	var err error
	config, err := pgxpool.ParseConfig(databaseURL)

	if err != nil {
		slog.Error("unable to parse postgresql dsn ", "error", err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		slog.Error("unable to create connection pool", "error", err)
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		slog.Error("unable to ping database", "error", err)
		pool.Close()
		return nil, err
	}

	slog.Info("successfully connected to postgresql database", "msg", "success")
	return pool, nil
}
