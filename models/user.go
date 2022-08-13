package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sync"
)

var (
	onceGetUser sync.Once
	uDoc        *mongo.Collection
)

//go:generate easytags $GOFILE json,xml,bson

type User struct {
	Id     int    `json:"id" xml:"id" bson:"id"`
	Name   string `json:"name" xml:"name" bson:"name"`
	Phone  string `json:"phone" xml:"phone" bson:"phone"`
	Email  string `json:"email" xml:"email" bson:"email"`
	Image  string `json:"image" xml:"image" bson:"image"`
	Gender string `json:"gender" xml:"gender" bson:"gender"`
	Rank   int    `json:"rank" xml:"rank" bson:"rank"`
}

func getUserDocument() *mongo.Collection {
	onceGetAccount.Do(func() {
		uDoc = db.Database("neworderfood").Collection("User")
	})
	return uDoc
}

func getCurrentUserDocument() (int, error) {
	i, err := getUserDocument().CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return 0, err
	}
	return int(i + 1), nil
}

func (this *User) UpdateUser() error {
	filter := bson.D{{"id", this.Id}}

	update := bson.D{
		{"$set", bson.D{
			{"name", this.Name},
			{"phone", this.Phone},
			{"email", this.Gender},
			{"image", this.Email},
			{"gender", this.Image},
		}},
	}

	_, err := getUserDocument().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/user.go:53")
		return err
	}

	return nil
}

func (this *User) GetUser() (User, error) {
	u := User{}
	filter := bson.D{{"id", this.Id}}
	err = getAccountDocument().FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/user.go:63")
	}
	return u, nil
}

func (this *User) UpRank() error {
	filter := bson.D{{"id", this.Id}}

	update := bson.D{
		{"$set", bson.D{
			{"Rank", this.Rank},
		}},
	}

	_, err := getUserDocument().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/user.go:83")
		return err
	}

	return nil
}
