package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Test connection
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	fmt.Println("Redis connected successfully")
}

// SetCache - Stores data in Redis with expiration time
func SetCache(key string, value string, expiration time.Duration) error {
	return RedisClient.Set(Ctx, key, value, expiration).Err()
}

// GetCache - Retrieves data from Redis
func GetCache(key string) (string, error) {
	return RedisClient.Get(Ctx, key).Result()
}

