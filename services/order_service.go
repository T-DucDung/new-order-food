package services

import (
	"errors"
	"new-order-food/models"
	"new-order-food/requests"
	"new-order-food/responses"
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

	err = order.PayOrder(order, lod, total)
	if err != nil {
		return err
	}

	total, err = GetTotal(strconv.Itoa(uid))
	if err != nil {
		return err
	}

	return Del(uid)
}

func CaculatorOrder(req []requests.RequestOrderDetail) ([]models.InsOrderDetail, float64, error) {
	lod := []models.InsOrderDetail{}
	orderDetail := models.InsOrderDetail{}
	p := models.Product{}
	var total float64
	total = 0

	for _, v := range req {
		orderDetail.ProductId = v.ProductId

		remaining, err := p.CheckRemaining(v.ProductId)
		if v.Quantity > remaining {
			return nil, -1, errors.New("not enough product, id: " + strconv.Itoa(v.ProductId))
		}
		orderDetail.Quantity = v.Quantity

		isSale, price, salePrice, err := p.GetPrice(v.ProductId)
		if err != nil || price == -1 || salePrice == -1 {
			return nil, 0, err
		}
		orderDetail.IsSale = isSale
		if isSale == true {
			orderDetail.Price = salePrice
		} else {
			orderDetail.Price = price
		}

		total = total + (price * float64(v.Quantity))
		lod = append(lod, orderDetail)
	}
	return lod, total, nil
}

func GetListOrder(uid string) ([]responses.OrderRes, error) {
	o := models.Order{}
	return o.GetListOrder(uid)
}

func GetListOrderForAdmin() ([]responses.OrderRes, error) {
	o := models.Order{}
	return o.GetListOrderForAdmin()
}

func UpDateOrder(id string) error {
	o := models.Order{}
	return o.UpdateOrder(id)
}

func GetTotal(id string) (float64, error) {
	o := models.Order{}
	return o.GetTotal(id)
}
