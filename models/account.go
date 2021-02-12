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

func (this *Account) RegisterForAdmin(req requests.RequestRegisterForAdmin) error {
	if req.Type == "user" {
		return this.Register(requests.RequestRegister{
			Username: req.Username,
			Pass:     req.Pass,
			Name:     req.Name,
			Phone:    req.Phone,
			Email:    req.Email,
			Image:    req.Image,
			Gender:   req.Gender,
		})
	} else if req.Type == "admin" {
		data, err := db.Prepare("INSERT INTO Admin(Name, Phone, Email, Image, Gender) VALUES(?, ?, ?, ?, ?);")
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
		_, err = data.Exec(req.Username, req.Pass, id, "admin")
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("account type not valid !")
}
