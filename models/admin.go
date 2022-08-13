package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"new-order-food/hash"
	"new-order-food/requests"
)

//go:generate easytags $GOFILE json,xml,bson

type Admin struct {
	Id     int    `json:"id" xml:"id" bson:"id"`
	Name   string `json:"name" xml:"name" bson:"name"`
	Phone  string `json:"phone" xml:"phone" bson:"phone"`
	Email  string `json:"email" xml:"email" bson:"email"`
	Image  string `json:"image" xml:"image" bson:"image"`
	Gender string `json:"gender" xml:"gender" bson:"gender"`
}

func (this *Admin) CreateAccount(req requests.RequestCreateAccount) error {
	if req.Type == "user" {
		acc := Account{}
		return acc.Register(requests.RequestRegister{
			Username: req.UserName,
			Pass:     req.Pass,
			Name:     req.Name,
			Phone:    req.Phone,
			Email:    req.Email,
			Image:    req.Image,
			Gender:   req.Gender,
		})
	} else if req.Type == "admin" {
		id, err := getCurrentUserDocument()
		if err != nil {
			log.Println(err.Error(), "err.Error() models/admin.go:34")
			return err
		}
		_, err = getUserDocument().InsertOne(context.TODO(), User{
			Id:     id,
			Name:   req.Name,
			Phone:  req.Phone,
			Email:  req.Email,
			Image:  req.Image,
			Gender: req.Gender,
			Rank:   1,
		})
		if err != nil {
			log.Println(err.Error(), "err.Error() models/admin.go:48")
			return err
		}
		hashP, _ := hash.Hash(req.Pass)
		_, err = getUserDocument().InsertOne(context.TODO(), Account{
			Username: req.UserName,
			Pass:     hashP,
			Id:       id,
			Type:     "admin",
			Status:   true,
		})
		if err != nil {
			log.Println(err.Error(), "err.Error() models/admin.go:60")
			return err
		}
	}
	return errors.New("account type not valid !")
}

func (this *Admin) GetAllUser(status string) ([]User, error) {
	results := make([]User, 0)

	b := bson.D{}
	if status != "" {
		b = bson.D{{"status", status}}
	}

	cur, err := getAccountDocument().Find(context.TODO(), b)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/admin.go:73")
		return []User{}, err
	}

	for cur.Next(context.TODO()) {
		var elem Admin
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/admin.go:80")
			return []User{}, err
		}

		u, _ := (&User{Id: elem.Id}).GetUser()

		results = append(results, u)
	}

	if err := cur.Err(); err != nil {
		log.Println(err.Error(), "err.Error() models/admin.go:89")
		return []User{}, err
	}

	return results, nil
}

func (this *Admin) UpdateStatus(req requests.RequestUpdateStatus) error {
	filter := bson.D{{"id", req.IdUser}}
	_, err := getAccountDocument().UpdateOne(context.TODO(), filter, bson.M{"$set": bson.M{"status": req.Status}})
	if err != nil {
		log.Println(err.Error(), "err.Error() models/admin.go:113")
		return err
	}

	return nil
}
