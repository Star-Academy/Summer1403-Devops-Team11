package helper

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	ctx    = context.Background()
	client = initRedisClient()
)

func initRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     GetEnv("REDIS_HOST", "redis:6379"),
		Password: "",
		DB:       0,
	})
}

func StoreResults[T any](key string, data []T) {
	jsonData, _ := json.Marshal(data)

	err := client.Set(ctx, key, jsonData, 0).Err()
	if err != nil {
		log.Fatalf("could not set data on redis: %v", err)
	}
}

func RetrieveResults[T any](key string) ([]T, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var retData []T
	err = json.Unmarshal([]byte(val), &retData)
	if err != nil {
		return nil, err
	}

	return retData, nil
}
