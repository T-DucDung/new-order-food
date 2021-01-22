package models

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
	err  error
)

func InitConnectDataBase() {
	once.Do(func() {
		db, err = sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/orderfood")
		if err != nil {
			log.Println("error connect database : ", err)
		} else {
			log.Println("====InitConnect====")
			log.Println(db)
			log.Println("===================")
		}
	})
}

func GetDatabase() *sql.DB {
	return db
}
