package services

import (
	"context"
	"errors"
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

func Get(key1 string, key2 string) (string, error) {
	r := New()
	val, err := r.Client.HGet(r.Context, key1, key2).Result()
	if err == redis.Nil {
		return "", errors.New("key not found")
	}

	if err != nil {
		panic(err)
	}
	return val, nil
}

func Set(key1 string, key2 string, value string) {
	r := New()
	err := r.Client.HSet(r.Context, key1, key2, value).Err()
	if err != nil {
		panic(err)
	}
}

func Increment(key1 string, key2 string) int {
	r := New()
	val, err := r.Client.HIncrBy(r.Context, key1, key2, 1).Result()
	if err != nil {
		panic(err)
	}
	return int(val)
}
