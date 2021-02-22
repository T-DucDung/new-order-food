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
	Username string `json:"username" xml:"username"`
	Pass     string `json:"pass" xml:"pass"`
	Id       int    `json:"id" xml:"id"`
	Type     string `json:"type" xml:"type"`
}

func (this *Account) Login(req requests.RequestLogin) (string, string, error) {
	a := Account{}
	err = db.QueryRow(queries.GetAccount(req.Username)).Scan(&a.Username, &a.Pass, &a.Id, &a.Type)
	if err != nil {
		return "", "", err
	}
	if a.Pass == req.Pass {
		data := "hehe" + a.Pass + strconv.FormatInt(time.Now().Unix(), 10)
		token, err := hash.Hash(data)
		if err != nil {
			return "", "", err
		}
		value := strconv.Itoa(a.Id) + ":" + a.Type
		_, err = Set(token, value, 2*time.Hour)
		if err != nil {
			log.Println("models/account.go:34 ", err)
			return "", "", err
		}

		return token, a.Type, nil
	}
	return "", "", errors.New("Login false ! Check the password")
}

func (this *Account) Register(req requests.RequestRegister) error {
	data, err := db.Prepare("INSERT INTO Users(Name, Phone, Email, Image, Gender) VALUES(?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	val, err := data.Exec(req.Name, req.Phone, req.Email, req.Image, req.Gender)
	if err != nil {
		return err
	}

	id, err := val.LastInsertId()

	data, err = db.Prepare("INSERT INTO Account(Username, Pass, Id, type) VALUES(?, ?, ?, ?);")
	if err != nil {
		return err
	}
	_, err = data.Exec(req.Username, req.Pass, id, "user")
	if err != nil {
		return err
	}
	return nil
}
