package models

import (
	"time"
)

func Set(key string, value interface{}, t time.Duration) (bool, error) {
	return rdb.SetNX(ctx, key, value, t).Result()
}

func Get(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}
