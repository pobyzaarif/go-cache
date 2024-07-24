package cache

import (
	"encoding/json"
	"fmt"
	"time"

	lru "github.com/hnlq715/golang-lru"
)

// MemoryARCCacheRepository ...
type MemoryARCCacheRepository struct {
	cache *lru.ARCCache
}

// NewMemoryARCCacheRepository new memory arc cache service
func NewMemoryARCCacheRepository(size int) (*MemoryARCCacheRepository, error) {
	cache, err := lru.NewARC(size)

	if err != nil {
		return nil, err
	}

	return &MemoryARCCacheRepository{
		cache,
	}, nil
}

// Set cache based on given key
func (memoryCache *MemoryARCCacheRepository) Set(key string, value interface{}, expiration time.Duration) (err error) {
	byteValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	strValue := string(byteValue)

	memoryCache.cache.AddEx(key, strValue, expiration)

	return nil
}

// Get cache base on given key
func (memoryCache *MemoryARCCacheRepository) Get(key string, data interface{}) (err error) {
	result, ok := memoryCache.cache.Get(key)
	if !ok {
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
func (memoryCache *MemoryARCCacheRepository) Delete(key string) {
	memoryCache.cache.Remove(key)
}
