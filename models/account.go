package models

import (
	"errors"
	"log"
	"new-order-food/hash"
	"new-order-food/queries"
	"new-order-food/requests"
	"strconv"
	"time"
)

type Account struct {
	Username string
	Pass     string
	Id       int
	Type     string
}

func (this *Account) Login(req requests.RequestLogin) (string, error) {
	a := Account{}
	err = db.QueryRow(queries.GetAccount(req.Username)).Scan(&a.Username, &a.Pass, &a.Id, &a.Type)
	if err != nil {
		return "", err
	}
	if a.Pass == req.Pass {
		data := "hehe" + a.Pass + strconv.FormatInt(time.Now().Unix(), 10)
		token, err := hash.Hash(data)
		if err != nil {
			return "", err
		}
		value := string(a.Id) + ":" + a.Type
		_, err = Set(token, value, 30*time.Second)
		if err != nil {
			log.Println("models/account.go:34 ", err)
			return "", err
		}

		return token, nil
	}
	return "", errors.New("Login false ! Check the password")
}
