package cache

import "time"

// Repository cache repository
type Repository interface {
	Set(key string, value interface{}, expiration time.Duration) (err error)

	Get(key string, data interface{}) (err error)

	Delete(key string)
}
