package cache_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	cache "github.com/pobyzaarif/go-cache"
	redis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryCache(t *testing.T) {
	staging := true

	var repo cache.Repository

	if staging {
		var errMemCache error
		repo, errMemCache = cache.NewMemoryARCCacheRepository(15) // increase 1 the size to avoid overcapacity of the cache because lru, lfu, and arc will remove unpopular keys automatically based on each algorithm
		assert.NoError(t, errMemCache)
	} else {
		// Create a new Redis client
		redisURL := os.Getenv("GOCACHE_REDIS_URL")
		if redisURL == "" {
			redisURL = "redis://localhost:6379/0?protocol=3"
		}
		opts, err := redis.ParseURL(redisURL)
		if err != nil {
			log.Fatalf("Could not connect to Redis: %v", err)
		}

		rdb := redis.NewClient(opts)

		// Ping the Redis server to check the connection
		_, err = rdb.Ping(context.TODO()).Result()
		if err != nil {
			log.Fatalf("Could not connect to Redis: %v", err)
		}

		repo = cache.NewRedisCacheRepository(rdb)
	}

	var test int
	err := repo.Set("test", 1, 5*time.Minute)
	assert.NoError(t, err)

	err = repo.Get("test", &test)
	assert.NoError(t, err)
	assert.Equal(t, 1, test)

	repo.Delete("test")

	var testB int
	err = repo.Get("test", &testB)
	assert.NoError(t, err)
	assert.Equal(t, 0, testB)
}
