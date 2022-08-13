package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math"
	"new-order-food/responses"
	"sync"
)

var (
	onceGetProduct sync.Once
	aProduct       *mongo.Collection
)

//go:generate easytags $GOFILE json,xml,bson

type Product struct {
	Id          int     `json:"id" xml:"id" bson:"id"`
	Name        string  `json:"name" xml:"name" bson:"name"`
	Image       string  `json:"image" xml:"image" bson:"image"`
	Price       float64 `json:"price" xml:"price" bson:"price"`
	IsSale      bool    `json:"is_sale" xml:"is_sale" bson:"is_sale"`
	Unit        string  `json:"unit" xml:"unit" bson:"unit"`
	Remaining   int     `json:"remaining" xml:"remaining" bson:"remaining"`
	SalePrice   float64 `json:"sale_price" xml:"sale_price" bson:"sale_price"`
	Description string  `json:"description" xml:"description" bson:"description"`
	Sold        int     `json:"sold" xml:"sold" bson:"sold"`
	CategoryId  int     `json:"category_id" xml:"category_id" bson:"category_id"`
	Rate1       int     `json:"rate1" xml:"rate1" bson:"rate1"`
	Rate2       int     `json:"rate2" xml:"rate2" bson:"rate2"`
	Rate3       int     `json:"rate3" xml:"rate3" bson:"rate3"`
	Rate4       int     `json:"rate4" xml:"rate4" bson:"rate4"`
	Rate5       int     `json:"rate5" xml:"rate5" bson:"rate5"`
	RateAvg     float64 `json:"rate_avg" xml:"rate_avg" bson:"rate_avg"`
}

func getProductDocument() *mongo.Collection {
	onceGetProduct.Do(func() {
		aProduct = db.Database("neworderfood").Collection("Product")
	})
	return aProduct
}

func getCurrentProductDocument() (int, error) {
	i, err := getProductDocument().CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return 0, err
	}
	return int(i + 1), nil
}

func (this *Product) CreateProduct() error {
	this.Id, err = getCurrentProductDocument()
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:58")
		return err
	}
	_, err = getProductDocument().InsertOne(context.TODO(), this)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:49")
		return err
	}
	return nil
}

func (this *Product) GetProduct(id int) (responses.ProductRes, error) {
	p := responses.ProductRes{}
	filter := bson.D{{"id", id}}
	err = getProductDocument().FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:74")
		return responses.ProductRes{}, err
	}

	return p, nil
}

func (this *Product) GetName() string {
	p := responses.ProductRes{}
	filter := bson.D{{"id", this.Id}}
	err = getProductDocument().FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:86")
		return ""
	}
	return p.Name
}

func (this *Product) GetListProduct(filter primitive.D) ([]responses.ProductRes, error) {
	lp := make([]responses.ProductRes, 0)

	cur, err := getProductDocument().Find(context.TODO(), filter)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:87")
		return []responses.ProductRes{}, err
	}

	for cur.Next(context.TODO()) {
		var elem responses.ProductRes
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/product.go:95")
			return []responses.ProductRes{}, err
		}

		lp = append(lp, elem)
	}

	if err := cur.Err(); err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:104")
		return []responses.ProductRes{}, err
	}

	return lp, nil
}

func (this *Product) UpDateProduct() error {
	filter := bson.D{{"id", this.Id}}

	update := bson.D{
		{"$set", bson.D{
			{"name", this.Name},
			{"image", this.Image},
			{"price", this.Price},
			{"is_sale", this.IsSale},
			{"unit", this.Unit},
			{"remaining", this.Remaining},
			{"sale_price", this.SalePrice},
			{"description", this.Description},
			{"sold", this.Sold},
			{"category_id", this.CategoryId},
			{"rate1", this.Rate1},
			{"rate2", this.Rate2},
			{"rate3", this.Rate3},
			{"rate4", this.Rate4},
			{"rate5", this.Rate5},
			{"rate_avg", this.RateAvg},
		}},
	}

	_, err = getProductDocument().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:136")
		return err
	}

	return nil
}

func (this *Product) CheckRemaining(pid int) (int, error) {
	p, err := this.GetProduct(pid)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:146")
		return 0, err
	}
	return p.Remaining, nil
}

func (this *Product) UpdateRemaining(pid int, total int, sold int) error {
	p := responses.ProductRes{}
	filter := bson.D{{"id", pid}}
	err = getProductDocument().FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:157")
		return err
	}

	update := bson.D{
		{"$set", bson.D{
			{"remaining", p.Remaining + total},
			{"sold", p.Sold + sold},
		}},
	}

	_, err = getProductDocument().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:170")
		return err
	}

	return nil
}

func (this *Product) UpdateSalePrice(pid string, salePrice float64, status bool) error {
	filter := bson.D{{"id", pid}}
	update := bson.D{
		{"$set", bson.D{
			{"sale_price", salePrice},
			{"is_sale", status},
		}},
	}

	_, err = getProductDocument().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:188")
		return err
	}

	return nil
}

func (this *Product) GetPrice(pid int) (bool, float64, float64, error) {
	p, err := this.GetProduct(pid)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:179")
		return false, 0, 0, err
	}
	return p.IsSale, p.Price, p.SalePrice, nil
}

func (this *Product) UpdateRate(pid int, last int, cur int) error {
	p := Product{}
	filter := bson.D{{"id", pid}}
	err = getProductDocument().FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:219")
		return err
	}
	cLast := 0
	nameLast := ""
	cCur := 0
	nameCur := ""
	update := bson.D{}

	switch last {
	case 0:
	case 1:
		cLast = p.Rate1
		nameLast = "rate1"
	case 2:
		cLast = p.Rate2
		nameLast = "rate2"
	case 3:
		cLast = p.Rate3
		nameLast = "rate3"
	case 4:
		cLast = p.Rate4
		nameLast = "rate4"
	case 5:
		cLast = p.Rate5
		nameLast = "rate5"
	}

	switch cur {
	case 1:
		cCur = p.Rate1
		nameCur = "rate1"
	case 2:
		cCur = p.Rate2
		nameCur = "rate2"
	case 3:
		cCur = p.Rate3
		nameCur = "rate3"
	case 4:
		cCur = p.Rate4
		nameCur = "rate4"
	case 5:
		cCur = p.Rate5
		nameCur = "rate5"
	}

	if last != 0 {
		update = bson.D{
			{"$set", bson.D{
				{nameLast, cLast - 1},
				{nameCur, cCur + 1},
			}},
		}

		_, err = getProductDocument().UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/product.go:251")
			return err
		}
	} else {
		update = bson.D{
			{"$set", bson.D{
				{nameCur, cCur + 1},
			}},
		}

		_, err = getProductDocument().UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/product.go:263")
			return err
		}
	}

	rate1, rate2, rate3, rate4, rate5 := this.GetRate(pid)

	total := math.Round((float64(rate1*1+rate2*2+rate3*3+rate4*4+rate5*5)/float64(rate1+rate2+rate3+rate4+rate5))*10) / 10

	update = bson.D{
		{"$set", bson.D{
			{"rate_avg", total},
		}},
	}

	_, err = getProductDocument().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:263")
		return err
	}
	return nil
}

func (this *Product) CheckExist(pid int) (bool, error) {
	p := Product{}
	filter := bson.D{{"id", pid}}
	err = getProductDocument().FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:314")
		return false, nil
	}
	return true, nil
}

func (this *Product) GetRate(id int) (int, int, int, int, int) {
	p := Product{}
	filter := bson.D{{"id", id}}
	err = getProductDocument().FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/product.go:75")
		return 0, 0, 0, 0, 0
	}

	return p.Rate1, p.Rate2, p.Rate3, p.Rate4, p.Rate5
}
