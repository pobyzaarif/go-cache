package cache_test

import (
	"testing"
	"time"

	cache "github.com/pobyzaarif/go-cache"
	"github.com/stretchr/testify/assert"
)

func TestMemoryCache(t *testing.T) {
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
	memCache, err := cache.NewMemoryARCCacheRepository(len(tc) + 1) // increase 1 the size to avoid overcapacity of the cache because lru, lfu, and arc will remove unpopular keys automatically based on each algorithm
	assert.NoError(t, err)

	var aa int
	err = memCache.Get("a", &aa)
	assert.NoError(t, err)
	assert.Equal(t, 0, aa)

	for k, v := range tc {
		err := memCache.Set(k, v, 5*time.Minute)
		assert.NoError(t, err)
	}

	// a test
	err = memCache.Get("a", &aa)
	assert.NoError(t, err)
	assert.Equal(t, tc["a"], aa)

	var aaa int
	memCache.Delete("a")
	memCache.Get("a", &aaa)
	assert.Equal(t, 0, aaa)

	// b test
	var bb float64
	err = memCache.Get("b", &bb)
	assert.NoError(t, err)
	assert.Equal(t, tc["b"], bb)

	// c test
	var cc string
	err = memCache.Get("c", &cc)
	assert.NoError(t, err)
	assert.Equal(t, tc["c"], cc)

	// d test
	var dd bool
	err = memCache.Get("d", &dd)
	assert.NoError(t, err)
	assert.Equal(t, tc["d"], dd)

	// f test
	var ff time.Time
	err = memCache.Get("f", &ff)
	assert.NoError(t, err)
	assert.WithinDuration(t, tc["f"].(time.Time), ff, time.Second)

	// g test
	var gg []string
	err = memCache.Get("g", &gg)
	assert.NoError(t, err)
	assert.Equal(t, tc["g"], gg)

	// h test
	var hh map[string]int
	err = memCache.Get("h", &hh)
	assert.NoError(t, err)
	assert.Equal(t, tc["h"], hh)

	// h test negative case
	var hhh map[int]string
	err = memCache.Get("h", &hhh)
	assert.Error(t, err)
	assert.NotEqual(t, tc["h"], hhh)
}
