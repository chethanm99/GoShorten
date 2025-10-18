package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	RDB *redis.Client
	Ctx = context.Background()
)

func Connect() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       0,
	})

	if _, err := RDB.Ping(Ctx).Result(); err != nil {
		return fmt.Errorf("redis connection failed: %w", err)
	}

	return nil
}

func IsDatabaseConnected() error {
	if RDB == nil {
		return fmt.Errorf("redis client not initialized")
	}
	_, err := RDB.Ping(Ctx).Result()
	return err
}
