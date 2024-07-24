package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

// RedisCacheRepository redis cache service
type RedisCacheRepository struct {
	client *redis.Client
}

// NewRedisCacheRepository new redis cache service
func NewRedisCacheRepository(client *redis.Client) *RedisCacheRepository {
	return &RedisCacheRepository{
		client,
	}
}

// Set set cache based on given key
func (redisCache *RedisCacheRepository) Set(key string, value interface{}, expiration time.Duration) error {
	byteValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	strValue := string(byteValue)

	return redisCache.client.Set(context.TODO(), key, strValue, expiration).Err()
}

// Get get cache base on given key
func (redisCache *RedisCacheRepository) Get(key string, data interface{}) (err error) {
	result, err := redisCache.client.Get(context.TODO(), key).Result()
	if err != nil {
		return nil
	}

	value := fmt.Sprintf("%v", result)
	if value != "" {
		err = json.Unmarshal([]byte(value), &data)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete cache base on given key
func (redisCache *RedisCacheRepository) Delete(key string) {
	redisCache.client.Expire(context.TODO(), key, 0)
}
