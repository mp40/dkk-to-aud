package cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

type CacheInterface interface {
	Set(key string, value string) (string, error)
	Get(key string) (string, error)
}

type RedisCache struct {
	client *redis.Client
}

var Cache RedisCache

func CreateRedisCache() {
	client := redis.NewClient(&redis.Options{
		Addr: "dkk_to_aud_redis:6379",
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Panic(err)
	}

	Cache.client = client
}

func (r *RedisCache) Set(key string, value string) (string, error) {
	_, err := r.client.Set(key, value, 6*time.Hour).Result()
	if err != nil {
		return "", err
	}

	return key, nil
}

func (r *RedisCache) Get(key string) (string, error) {
	value, err := r.client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}
