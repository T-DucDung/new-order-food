package services

import (
	"errors"
	"new-order-food/models"
	"new-order-food/responses"
)

func Set(uid, pid, quantity int) error {
	if quantity <= 0 {
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
	cart := models.Cart{
		Userid:    uid,
		ProductId: pid,
		Quantity:  quantity,
	}
	return cart.UpdateItem()
}
