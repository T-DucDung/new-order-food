package models

import (
	"context"
	"database/sql"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db    *sql.DB
	onceM sync.Once
	onceR sync.Once
	err   error
	rdb   *redis.Client
)

var ctx = context.Background()

func InitConnectDataBase() {
	onceM.Do(func() {
		db, err = sql.Open("mysql", "root:Dung13524685@tcp(127.0.0.1:3306)/neworderfood")
		if err != nil {
			log.Println("error connect database : ", err)
		} else {
			log.Println("====InitConnectMySql====")
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
