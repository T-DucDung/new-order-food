package models

import (
	"strconv"
	"time"
)

type Order struct {
	Id            int `json:"id" xml:"id"`
	UserId        int `json:"user_id" xml:"user_id"`
	lastUpdate    int64
	Number        string        `json:"number" xml:"number"`
	Address       string        `json:"address" xml:"address"`
	Name          string        `json:"name" xml:"name"`
	Total         float32       `json:"total" xml:"total"`
	CurrentStatus string        `json:"current_status" xml:"current_status"`
	Detail        []OrderDetail `json:"detail" xml:"detail"`
}

type OrderDetail struct {
	ProductId int     `json:"product_id" xml:"product_id"`
	IsSale    bool    `json:"is_sale" xml:"is_sale"`
	Price     float32 `json:"price" xml:"price"`
	Quantity  int     `json:"quantity" xml:"quantity"`
}

func (this *Order) PayOrder(order Order, lod []OrderDetail, total float32) error {
	data, err := db.Prepare("insert into `Order` (UserId,Name,Phone,Address,Total,CurrentStatus,LastUpDate) values VALUES(?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	val, err := data.Exec(order.UserId, order.Name, order.Number, order.Address, total, "", time.Now().Unix())
	if err != nil {
		return err
	}

	id, err := val.LastInsertId()
	p := Product{}

	for _, item := range lod {
		data, err = db.Prepare("insert into OrderDetail (OrderId,ProductId,IsSale,Price,Quantity) VALUES(?, ?, ?, ?, ?);")
		if err != nil {
			return err
		}
		_, err = data.Exec(id, item.ProductId, item.IsSale, item.Price, item.Quantity)
		if err != nil {
			return err
		}
		err = p.UpdateRemaining(strconv.Itoa(item.ProductId), item.Quantity)
		if err != nil {
			return err
		}
	}

	return nil
}
