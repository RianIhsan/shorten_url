package cache

import "time"

type RedisCache interface {
	SetRdsShortURL(key string, ttl time.Duration, data any) error
}
