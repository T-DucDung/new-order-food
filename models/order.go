package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"new-order-food/responses"
	"strconv"
	"sync"
	"time"
)

var mapStatus = map[int]string{
	1: "Waiting for shipment",
	2: "Done",
}

//go:generate easytags $GOFILE json,xml,bson

var (
	onceGetOrder sync.Once
	aOrder       *mongo.Collection
)

func getOrderDocument() *mongo.Collection {
	onceGetOrder.Do(func() {
		aOrder = db.Database("neworderfood").Collection("Order")
	})
	return aOrder
}

func getCurrentOrderDocument() (int, error) {
	i, err := getOrderDocument().CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err.Error(), "err.Error() models/order.go:36")
		return 0, err
	}
	return int(i + 1), nil
}

var (
	onceGetOrderD sync.Once
	aOrderDetail       *mongo.Collection
)

func getOrderDetailDocument() *mongo.Collection {
	onceGetOrderD.Do(func() {
		aOrderDetail = db.Database("neworderfood").Collection("OrderDetail")
	})
	return aOrderDetail
}

type Order struct {
	Id            int     `json:"id" xml:"id" bson:"id"`
	UserId        int     `json:"user_id" xml:"user_id" bson:"user_id"`
	LastUpdate    int64   `json:"last_update" xml:"last_update" bson:"last_update"`
	Number        string  `json:"number" xml:"number" bson:"number"`
	Address       string  `json:"address" xml:"address" bson:"address"`
	Name          string  `json:"name" xml:"name" bson:"name"`
	Total         float64 `json:"total" xml:"total" bson:"total"`
	CurrentStatus int     `json:"current_status" xml:"current_status" bson:"current_status"`
}

type InsOrderDetail struct {
	OrderId   int     `json:"order_id" xml:"order_id" bson:"order_id"`
	ProductId int     `json:"product_id" xml:"product_id" bson:"product_id"`
	IsSale    bool    `json:"is_sale" xml:"is_sale" bson:"is_sale"`
	Price     float64 `json:"price" xml:"price" bson:"price"`
	Quantity  int     `json:"quantity" xml:"quantity" bson:"quantity"`
}

func (this *Order) PayOrder(order Order, lod []InsOrderDetail, total float64) error {
	id, err := getCurrentOrderDocument()
	if err != nil {
		log.Println(err.Error(), "err.Error() models/order.go:62")
		return err
	}
	_, err = getOrderDocument().InsertOne(context.TODO(), Order{
		Id:            id,
		UserId:        order.UserId,
		LastUpdate:    time.Now().Unix(),
		Number:        order.Number,
		Address:       order.Address,
		Name:          order.Name,
		Total:         total,
		CurrentStatus: 1,
	})
	if err != nil {
		log.Println(err.Error(), "err.Error() models/order.go:76")
		return err
	}

	for _, item := range lod {
		log.Println(item, "item models/order.go:95")
		_, err = getOrderDetailDocument().InsertOne(context.TODO(), InsOrderDetail{
			OrderId:   id,
			ProductId: item.ProductId,
			IsSale:    item.IsSale,
			Price:     item.Price,
			Quantity:  item.Quantity,
		})
		if err != nil {
			log.Println(err.Error(), "err.Error() models/order.go:111")
			return err
		}

		err = (&Product{}).UpdateRemaining(item.ProductId, -item.Quantity, item.Quantity)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *Order) GetListOrder(uid string) ([]responses.OrderRes, error) {
	results := make([]responses.OrderRes, 0)

	b := bson.D{{"user_id", uid}}

	cur, err := getOrderDocument().Find(context.TODO(), b)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/order.go:129")
		return []responses.OrderRes{}, err
	}

	for cur.Next(context.TODO()) {
		var elem Order
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/order.go:137")
			return []responses.OrderRes{}, err
		}

		u, err := getOrderDetail(elem.Id)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/order.go:143")
			continue
		}

		results = append(results, responses.OrderRes{
			Id:            elem.Id,
			Number:        elem.Number,
			Address:       elem.Address,
			Name:          elem.Name,
			LastUpdate:    elem.LastUpdate,
			Total:         elem.Total,
			CurrentStatus: strconv.Itoa(elem.CurrentStatus),
			Detail:        u,
		})
	}

	if err := cur.Err(); err != nil {
		log.Println(err.Error(), "err.Error() models/order.go:160")
		return []responses.OrderRes{}, err
	}

	return results, nil
}

func getOrderDetail(id int) ([]responses.OrderDetailRes, error) {
	results := make([]responses.OrderDetailRes, 0)

	b := bson.D{{"order_id", id}}

	cur, err := getOrderDetailDocument().Find(context.TODO(), b)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/order.go:174")
		return []responses.OrderDetailRes{}, err
	}

	for cur.Next(context.TODO()) {
		var elem InsOrderDetail
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/order.go:182")
			return []responses.OrderDetailRes{}, err
		}

		results = append(results, responses.OrderDetailRes{
			Name:     (&Product{Id: elem.ProductId}).GetName(),
			IsSale:   elem.IsSale,
			Price:    elem.Price,
			Quantity: elem.Quantity,
		})
	}

	if err := cur.Err(); err != nil {
		log.Println(err.Error(), "err.Error() models/order.go:195")
		return []responses.OrderDetailRes{}, err
	}

	return results, nil
}

func (this *Order) GetListOrderForAdmin() ([]responses.OrderRes, error) {
	results := make([]responses.OrderRes, 0)

	cur, err := getOrderDocument().Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err.Error(), "err.Error() models/order.go:207")
		return []responses.OrderRes{}, err
	}

	for cur.Next(context.TODO()) {
		var elem Order
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/order.go:215")
			return []responses.OrderRes{}, err
		}

		u, err := getOrderDetail(elem.Id)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/order.go:221")
			continue
		}

		results = append(results, responses.OrderRes{
			Id:            elem.Id,
			Number:        elem.Number,
			Address:       elem.Address,
			Name:          elem.Name,
			LastUpdate:    elem.LastUpdate,
			Total:         elem.Total,
			CurrentStatus: strconv.Itoa(elem.CurrentStatus),
			Detail:        u,
		})
	}

	if err := cur.Err(); err != nil {
		log.Println(err.Error(), "err.Error() models/order.go:238")
		return []responses.OrderRes{}, err
	}

	return results, nil
}

func (this *Order) UpdateOrder(id string) error {
	filter := bson.D{{"id", id}}

	update := bson.D{
		{"$set", bson.D{
			{"current_status", 2},
		}},
	}

	_, err := getOrderDocument().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/order.go:256")
		return err
	}
	return nil
}

func (this *Order) GetTotal(id string) (float64, error) {
	total, err := getOrderDocument().CountDocuments(context.TODO(), bson.M{"user_id": id})
	if err != nil {
		log.Println(err.Error(), "err.Error() models/order.go:264")
		return 0, err
	}
	return float64(total), nil
}
