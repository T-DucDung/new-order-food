package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	db    *mongo.Client
	onceM sync.Once
	onceR sync.Once
	err   error
	rdb   *redis.Client
)

var ctx = context.Background()

func InitConnectDataBase() {
	onceM.Do(func() {
		db, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Println("error connect database : ", err)
		} else {
			log.Println("====InitConnectMongoDb====")
			log.Println(db)
			log.Println("========================")
		}
	})
}

func InitRedisClient() {
	onceR.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		})
		pong, err := rdb.Ping(ctx).Result()
		log.Println("====InitConnectRedis====")
		log.Println(pong, err)
		log.Println("========================")
	})
}
