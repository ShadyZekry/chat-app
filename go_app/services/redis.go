package services

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	Client  *redis.Client
	Context context.Context
}

var redisService *RedisService
var once sync.Once

func New() *RedisService {
	once.Do(func() {
		redisService = &RedisService{
			Client: redis.NewClient(&redis.Options{
				Addr:     "redis:6379",
				Password: "",
				DB:       0,
			}),
			Context: context.Background(),
		}
	})
	return redisService
}

func Get(key string) string {
	r := New()
	val, err := r.Client.Get(r.Context, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func Set(key string, value string) string {
	r := New()
	err := r.Client.Set(r.Context, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	return value
}

func Increment(key string) int {
	r := New()
	val, err := r.Client.Incr(r.Context, key).Result()
	if err != nil {
		panic(err)
	}
	return int(val)
}
