package models

import (
	"errors"
	"new-order-food/requests"
)

type Admin struct {
	Id     int    `json:"id" xml:"id"`
	Name   string `json:"name" xml:"name"`
	Phone  string `json:"phone" xml:"phone"`
	Email  string `json:"email" xml:"email"`
	Image  string `json:"image" xml:"image"`
	Gender string `json:"gender" xml:"gender"`
}

func (this *Admin) CreateAccount(req requests.RequestCreateAccount) error {
	if req.Type == "user" {
		acc := Account{}
		return acc.Register(requests.RequestRegister{
			Username: req.Username,
			Pass:     req.Pass,
			Name:     req.Name,
			Phone:    req.Phone,
			Email:    req.Email,
			Image:    req.Image,
			Gender:   req.Gender,
		})
	} else if req.Type == "admin" {
		data, err := db.Prepare("INSERT INTO User(Name, Phone, Email, Image, Gender, Rank) VALUES(?, ?, ?, ?, ?, ?);")
		if err != nil {
			return err
		}
		val, err := data.Exec(req.Name, req.Phone, req.Email, req.Image, req.Gender, "0")
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
