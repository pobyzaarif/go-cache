package cache_test

import (
	"context"
	"log"
	"testing"
	"time"

	cache "github.com/pobyzaarif/go-cache"
	redis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisCache(t *testing.T) {
	timeNow := time.Now()
	tc := map[string]interface{}{
		"a": 1,
		"b": 2.5,
		"c": "3",
		"d": true,
		"f": timeNow,
		"g": []string{"1", "2", "3"},
		"h": map[string]int{"1": 2, "3": 4},
		"i": 1,
	}

	// Create a new Redis client
	url := "redis://localhost:6379/0?protocol=3"
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	rdb := redis.NewClient(opts)

	// Ping the Redis server to check the connection
	_, err = rdb.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	redisCache := cache.NewRedisCacheRepository(rdb) // increase 1 the size to avoid overcapacity of the cache because lru, lfu, and arc will remove unpopular keys automatically based on each algorithm
	assert.NoError(t, err)

	var aa int
	err = redisCache.Get("a", &aa)
	assert.NoError(t, err)
	assert.Equal(t, 0, aa)

	for k, v := range tc {
		err := redisCache.Set(k, v, 5*time.Minute)
		assert.NoError(t, err)
	}

	// a test
	err = redisCache.Get("a", &aa)
	assert.NoError(t, err)
	assert.Equal(t, tc["a"], aa)

	var aaa int
	redisCache.Delete("a")
	redisCache.Get("a", &aaa)
	assert.Equal(t, 0, aaa)

	// b test
	var bb float64
	err = redisCache.Get("b", &bb)
	assert.NoError(t, err)
	assert.Equal(t, tc["b"], bb)

	// c test
	var cc string
	err = redisCache.Get("c", &cc)
	assert.NoError(t, err)
	assert.Equal(t, tc["c"], cc)

	// d test
	var dd bool
	err = redisCache.Get("d", &dd)
	assert.NoError(t, err)
	assert.Equal(t, tc["d"], dd)

	// f test
	var ff time.Time
	err = redisCache.Get("f", &ff)
	assert.NoError(t, err)
	assert.WithinDuration(t, tc["f"].(time.Time), ff, time.Second)

	// g test
	var gg []string
	err = redisCache.Get("g", &gg)
	assert.NoError(t, err)
	assert.Equal(t, tc["g"], gg)

	// h test
	var hh map[string]int
	err = redisCache.Get("h", &hh)
	assert.NoError(t, err)
	assert.Equal(t, tc["h"], hh)

	// h test negative case
	var hhh map[int]string
	err = redisCache.Get("h", &hhh)
	assert.Error(t, err)
	assert.NotEqual(t, tc["h"], hhh)
}
