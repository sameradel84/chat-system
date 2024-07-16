package database

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

// RedisClient is the global Redis client instance.
var RedisClient *redis.Client

// ctx is the global context used for Redis operations.
var ctx = context.Background()

// RedisConfig holds the configuration details required to connect to Redis.
type RedisConfig struct {
	Address string `json:"address"`
}

// RedisNewConn returns the active Redis client instance.
func RedisNewConn() *redis.Client {
	return RedisClient
}

// InitRedis initializes a connection to Redis using configuration from redis_config.json.
func InitRedis() {
	configFile := "config/redis_config.json"
	configData, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Failed to read Redis config file: %v", err)
	}

	var config RedisConfig
	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Failed to parse Redis config: %v", err)
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr: config.Address,
	})

	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")
}

// SetCache sets a key-value pair in Redis cache.
func SetCache(key string, value interface{}) error {
	return RedisClient.Set(ctx, key, value, 0).Err()
}

// GetCache retrieves a value from Redis cache based on the provided key.
func GetCache(key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

// InvalidateCache deletes one or more keys from Redis cache.
func InvalidateCache(keys ...string) error {
	return RedisClient.Del(ctx, keys...).Err()
}
