package services

import (
	"errors"
	"new-order-food/models"
	"new-order-food/requests"
	"strconv"
)

func PayOrder(req requests.RequestOrder, uid int) error {
	order := models.Order{
		UserId:  uid,
		Number:  req.Number,
		Address: req.Address,
		Name:    req.Name,
	}
	lod, total, err := CaculatorOrder(req.Detail)
	if err != nil || total == -1 {
		return err
	}
	return order.PayOrder(order, lod, total)
}

func CaculatorOrder(req []requests.RequestOrderDetail) ([]models.OrderDetail, float32, error) {
	lod := []models.OrderDetail{}
	orderDetail := models.OrderDetail{}
	p := models.Product{}
	var total float32
	total = 0

	for _, v := range req {
		orderDetail.ProductId = v.ProductId

		remaining, err := p.CheckRemaining(strconv.Itoa(v.ProductId))
		if v.Quantity < remaining {
			return nil, -1, errors.New("not enough product, id: " + strconv.Itoa(v.ProductId))
		}
		orderDetail.Quantity = v.Quantity

		isSale, price, err := p.GetPrice(strconv.Itoa(v.ProductId))
		if err != nil || price == -1 {
			return nil, 0, err
		}
		orderDetail.IsSale = isSale
		orderDetail.Price = price

		total = total + (price * float32(v.Quantity))
		lod = append(lod, orderDetail)
	}
	return lod, total, nil
}
