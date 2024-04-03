package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/RianIhsan/shorten_url/config"
	"github.com/RianIhsan/shorten_url/helper/cache"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type redisCacheRepository struct {
	rdb *redis.Client
}

func NewRedisClient(cnf config.APPConfig) cache.RedisCache {
	newConn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cnf.RdsConf.RdsHost, cnf.RdsConf.RdsPort),
		Password: cnf.RdsConf.RdsPass,
		DB:       0,
	})
	logrus.Info("Redis connected successfully")

	return &redisCacheRepository{
		rdb: newConn,
	}
}

func (r redisCacheRepository) SetRdsShortURL(key string, ttl time.Duration, data any) error {
	return r.rdb.Set(context.Background(), key, data, ttl).Err()
}

