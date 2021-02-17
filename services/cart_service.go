package services

import (
	"errors"
	"new-order-food/models"
	"new-order-food/responses"
	"strconv"
)

func Set(uid, pid, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity error")
	}
	p := models.Product{}
	total, err := p.CheckRemaining(strconv.Itoa(pid))
	if err != nil {
		return err
	}
	if quantity > total {
		return errors.New("quantity error")
	}

	cart := models.Cart{
		Userid:    uid,
		ProductId: pid,
		Quantity:  quantity,
	}
	return cart.Set()
}

func Get(uid int) ([]responses.CartRes, error) {
	cart := models.Cart{Userid: uid}
	return cart.Get()
}

func RemoveItem(uid, pid int) error {
	cart := models.Cart{
		Userid:    uid,
		ProductId: pid,
	}
	return cart.RemoveItem()
}

func Del(uid int) error {
	cart := models.Cart{Userid: uid}
	return cart.Del()
}

func UpdateItem(uid, pid, quantity int) error {
	if quantity == 0 {
		return RemoveItem(uid, pid)
	} else if quantity < 0 {
		return errors.New("quantity error")
	}
	p := models.Product{}
	total, err := p.CheckRemaining(strconv.Itoa(pid))
	if err != nil {
		return err
	}
	if quantity > total {
		return errors.New("quantity error")
	}

	cart := models.Cart{
		Userid:    uid,
		ProductId: pid,
		Quantity:  quantity,
	}
	return cart.UpdateItem()
}
