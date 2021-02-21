package models

import (
	"strconv"
	"time"
)

type Order struct {
	Id            int
	UserId        int
	lastUpdate    int64
	Number        string
	Address       string
	Name          string
	Total         float32
	CurrentStatus string
	Detail        []OrderDetail
}

type OrderDetail struct {
	ProductId int
	IsSale    bool
	Price     float32
	Quantity  int
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
