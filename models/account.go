package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"new-order-food/hash"
	"new-order-food/requests"
	"strconv"
	"sync"
	"time"
)

var (
	onceGetAccount sync.Once
	aDoc           *mongo.Collection
)

//go:generate easytags $GOFILE json,xml,bson

type Account struct {
	Username string `json:"username" xml:"username" bson:"username"`
	Pass     string `json:"pass" xml:"pass" bson:"pass"`
	Id       int    `json:"id" xml:"id" bson:"id"`
	Type     string `json:"type" xml:"type" bson:"type"`
	Status   bool   `json:"status" xml:"status" bson:"status"`
}

func getAccountDocument() *mongo.Collection {
	onceGetAccount.Do(func() {
		aDoc = db.Database("neworderfood").Collection("Account")
	})
	return aDoc
}

func (this *Account) Login(req requests.RequestLogin) (string, string, error) {
	a := Account{}
	filter := bson.D{{"username", req.Username}}
	err = getAccountDocument().FindOne(context.TODO(), filter).Decode(&a)
	if err != nil {
		return "", "", err
	}
	if a.Status == false {
		return "", "", errors.New("account is deactive")
	}
	hashP, _ := hash.Hash(req.Pass)
	//log.Println(hashP)
	if hashP == a.Pass {
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
	id, err := getCurrentUserDocument()
	if err != nil {
		log.Println(err.Error(), "err.Error() models/account.go:70")
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
		log.Println(err.Error(), "err.Error() models/account.go:83")
		return err
	}

	hashP, _ := hash.Hash(req.Pass)
	_, err = getUserDocument().InsertOne(context.TODO(), Account{
		Username: req.Username,
		Pass:     hashP,
		Id:       id,
		Type:     "user",
		Status:   true,
	})
	if err != nil {
		log.Println(err.Error(), "err.Error() models/account.go:103")
		return err
	}

	return nil
}
