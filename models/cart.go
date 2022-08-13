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
	onceGetCart sync.Once
	aCart       *mongo.Collection
)

//go:generate easytags $GOFILE json,xml,bson

type Cart struct {
	Userid    int `json:"userid" xml:"userid" bson:"userid"`
	ProductId int `json:"product_id" xml:"product_id" bson:"product_id"`
	Quantity  int `json:"quantity" xml:"quantity" bson:"quantity"`
}

func getCartDocument() *mongo.Collection {
	onceGetCart.Do(func() {
		aCart = db.Database("neworderfood").Collection("Cart")
	})
	return aCart
}

func (this *Cart) Set() error {
	_, err := getCartDocument().InsertOne(context.TODO(), this)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/cart.go:34")
	}

	return nil
}

func (this *Cart) Get() ([]responses.CartRes, error) {
	lc := []responses.CartRes{}

	b := bson.D{{"userid", this.Userid}}

	cur, err := getCartDocument().Find(context.TODO(), b)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/cart.go:50")
		return []responses.CartRes{}, err
	}

	for cur.Next(context.TODO()) {
		var elem Cart
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/cart.go:58")
			return []responses.CartRes{}, err
		}

		u, _ := (&Product{}).GetProduct(elem.ProductId)

		lc = append(lc, responses.CartRes{
			Id:        u.Id,
			Name:      u.Name,
			Quantity:  elem.Quantity,
			Price:     u.Price,
			IsSale:    u.IsSale,
			SalePrice: u.SalePrice,
		})
	}

	if err := cur.Err(); err != nil {
		log.Println(err.Error(), "err.Error() models/cart.go:68")
		return []responses.CartRes{}, err
	}

	return lc, nil
}

func (this *Cart) Del() error {
	filter := bson.D{{"userid", this.Userid}}

	d, err := getCartDocument().DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/cart.go:85")
		return err
	}
	log.Println(d.DeletedCount, "d.DeletedCount models/cart.go:88")

	return nil
}

func (this *Cart) RemoveItem() error {
	filter := bson.D{{"userid", this.Userid}, {"product_id", this.ProductId}}

	_, err = getCartDocument().DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/cart.go:97")
	}

	return nil
}

func (this *Cart) UpdateItem() error {
	filter := bson.D{{"userid", this.Userid}, {"product_id", this.ProductId}}

	update := bson.D{
		{"$set", bson.D{
			{"quantity", this.Quantity},
		}},
	}

	_, err = getCartDocument().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/cart.go:115")
		return err
	}
	return nil
}

func (this *Cart) CheckExist() (bool, error) {
	c := Cart{}
	filter := bson.D{{"userid", this.Userid}, {"product_id", this.ProductId}}
	err = getCartDocument().FindOne(context.TODO(), filter).Decode(&c)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/cart.go:126")
		return false, nil
	}
	return true, nil
}
