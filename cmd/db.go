package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DBPool *pgxpool.Pool

func InitDB() (*pgxpool.Pool, error) {
	connStr := os.Getenv("DATABASE_URL")
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		// TODO: Setup logging when this happens (FATAL)
		return nil, fmt.Errorf("Failed to parse connection string: %w", err)
	}

	// Configure Pool Settings
	poolConfig.MaxConns = 100
	poolConfig.MinConns = 10
	poolConfig.MaxConnLifetime = time.Hour / 2

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to create connection pool: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pool.Acquire(ctx) // Acquire a connection to test
	if err != nil {
		// TODO: Setup logging when this happens (FATAL)
		return nil, fmt.Errorf("Failed to acquire connection from pool: %w", err)
	}
	defer conn.Release()

	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Database connection test failed: %w", err)
	}
	fmt.Println("Database connection test successful.")
	return pool, nil
}
