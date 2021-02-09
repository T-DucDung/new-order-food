package redis

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	rdb *redis.Client
}

var (
	Client RedisClient
)

func InitRedisClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	Client = RedisClient{rdb: rdb}
}

func (this RedisClient) Set(key, value string, time time.Duration) error {
	err := this.rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (this RedisClient) Get(key string) (interface{}, error) {
	val, err := this.rdb.Get(ctx, "key").Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
