package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr            string
	Password        string
	DialTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	PoolSize        int
	MinIdleConns    int
	PoolTimeout     time.Duration
	ConnMaxIdleTime time.Duration
}

func NewRedisClient(ctx context.Context, config RedisConfig) (*redis.Client, error) {
	options := &redis.Options{
		Addr:            config.Addr,
		Password:        config.Password,
		DialTimeout:     config.DialTimeout,
		ReadTimeout:     config.ReadTimeout,
		WriteTimeout:    config.WriteTimeout,
		PoolSize:        config.PoolSize,
		MinIdleConns:    config.MinIdleConns,
		PoolTimeout:     config.PoolTimeout,
		ConnMaxIdleTime: config.ConnMaxIdleTime,
	}

	rdb := redis.NewClient(options)

	// Test connection using PING
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("Failed to connect to Redis: %w", err)
	}

	// TODO: Change this into a logging statement
	fmt.Println("Successfully connected to Redis.")
	return rdb, nil
}
