package models

import (
	"new-order-food/queries"
	"new-order-food/responses"
	"strconv"
	"time"
)

var mapStatus = map[int]string{
	1 : "Waiting for shipment",
	2 : "Done",
}

type Order struct {
	Id            int           `json:"id" xml:"id"`
	UserId        int           `json:"user_id" xml:"user_id"`
	LastUpdate    int64         `json:"last_update" xml:"last_update"`
	Number        string        `json:"number" xml:"number"`
	Address       string        `json:"address" xml:"address"`
	Name          string        `json:"name" xml:"name"`
	Total         float32       `json:"total" xml:"total"`
	CurrentStatus int           `json:"current_status" xml:"current_status"`
	Detail        []OrderDetail `json:"detail" xml:"detail"`
}

type OrderDetail struct {
	ProductId int     `json:"product_id" xml:"product_id"`
	IsSale    bool    `json:"is_sale" xml:"is_sale"`
	Price     float32 `json:"price" xml:"price"`
	Quantity  int     `json:"quantity" xml:"quantity"`
}

func (this *Order) PayOrder(order Order, lod []OrderDetail, total float32) error {
	data, err := db.Prepare("insert into `Order` (UserId,Name,Phone,Address,Total,CurrentStatus,LastUpDate) VALUES(?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	val, err := data.Exec(order.UserId, order.Name, order.Number, order.Address, total, 1, time.Now().Unix())
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
		err = p.UpdateRemaining(strconv.Itoa(item.ProductId), -item.Quantity, item.Quantity)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *Order) GetListOrder(uid string) ([]responses.OrderRes, error) {
	lo := []responses.OrderRes{}

	results, err := db.Query(queries.GetListOrder(uid))
	if err != nil {
		return nil, err
	}

	for results.Next() {
		o := responses.OrderRes{}
		status := 0
		err = results.Scan(&o.Id, &o.Name, &o.Number, &o.Total, &status, &o.LastUpdate)
		if err != nil {
			return nil, err
		}
		o.CurrentStatus = mapStatus[status]
		lo = append(lo, o)
	}

	for index, item := range lo {
		lod, err := getOrderDetail(strconv.Itoa(item.Id))
		if err != nil {
			return nil, err
		}
		lo[index].Detail = lod
	}

	return lo, nil
}

func getOrderDetail(id string) ([]responses.OrderDetailRes, error){
	lod := []responses.OrderDetailRes{}

	results, err := db.Query(queries.GetListOrderDetail(id))
	if err != nil {
		return nil, err
	}

	for results.Next() {
		o := responses.OrderDetailRes{}
		err = results.Scan(&o.Name, &o.IsSale, &o.Price, &o.Quantity)
		if err != nil {
			return nil, err
		}
		lod = append(lod, o)
	}

	return lod, nil
}

func (this *Order) GetListOrderForAdmin() ([]responses.OrderRes, error) {
	lo := []responses.OrderRes{}

	results, err := db.Query(queries.GetListOrderForAdmin())
	if err != nil {
		return nil, err
	}

	for results.Next() {
		o := responses.OrderRes{}
		status := 0
		err = results.Scan(&o.Id, &o.Name, &o.Number, &o.Address, &o.Total, &status, &o.LastUpdate)
		if err != nil {
			return nil, err
		}
		o.CurrentStatus = mapStatus[status]
		lo = append(lo, o)
	}

	for index, item := range lo {
		lod, err := getOrderDetail(strconv.Itoa(item.Id))
		if err != nil {
			return nil, err
		}
		lo[index].Detail = lod
	}

	return lo, nil
}

func (this *Order) UpdateOrder(id string) error {
	data, err := db.Prepare("update `Order` set CurrentStatus = 2 where id = ?")
	if err != nil {
		return err
	}
	_, err = data.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
