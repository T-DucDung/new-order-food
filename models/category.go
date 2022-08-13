package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"new-order-food/responses"
	"sync"
)

var (
	onceGetCate sync.Once
	aCate       *mongo.Collection
)

//go:generate easytags $GOFILE json,xml,bson

type Category struct {
	CategoryId int    `json:"category_id" xml:"category_id" bson:"category_id"`
	Name       string `json:"name" xml:"name" bson:"name"`
}

func getCateDocument() *mongo.Collection {
	onceGetCate.Do(func() {
		aCate = db.Database("neworderfood").Collection("Category")
	})
	return aCate
}

func getCurrentCateDocument() (int, error) {
	i, err := getCateDocument().CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return 0, err
	}
	return int(i + 1), nil
}

func (this *Category) GetListCategory() ([]responses.CategoryRes, error) {
	results := make([]responses.CategoryRes, 0)

	b := bson.D{}

	cur, err := getCateDocument().Find(context.TODO(), b)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/category.go:45")
		return []responses.CategoryRes{}, err
	}

	for cur.Next(context.TODO()) {
		var elem Category
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/category.go:53")
			return []responses.CategoryRes{}, err
		}

		count, err := getProductDocument().CountDocuments(context.TODO(), bson.M{"category_id": elem.CategoryId})
		if err != nil {
			log.Println(err.Error(), "err.Error() models/category.go:59")
			return []responses.CategoryRes{}, err
		}

		results = append(results, responses.CategoryRes{
			Id:    elem.CategoryId,
			Name:  elem.Name,
			Total: int(count),
		})
	}

	if err := cur.Err(); err != nil {
		log.Println(err.Error(), "err.Error() models/category.go:63")
		return []responses.CategoryRes{}, err
	}

	return results, nil
}

func (this *Category) UpdateCategory() error {
	filter := bson.D{{"category_id", this.CategoryId}}

	update := bson.D{
		{"$set", bson.D{
			{"Name", this.Name},
		}},
	}

	_, err = getCateDocument().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/category.go:88")
		return err
	}

	return nil
}

func (this *Category) CreateCategory() error {
	id, err := getCurrentCateDocument()
	if err != nil {
		log.Println(err.Error(), "err.Error() models/category.go:93")
		return err
	}
	this.CategoryId = id
	_, err = getCateDocument().InsertOne(context.TODO(), this)
	if err != nil {
		log.Println(err.Error(), "err models/category.go:99")
		return err
	}
	return nil
}
