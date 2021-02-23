package models

import (
	"new-order-food/queries"
	"new-order-food/responses"
	"strconv"
)

type Cart struct {
	Userid    int `json:"userid" xml:"userid"`
	ProductId int `json:"product_id" xml:"product_id"`
	Quantity  int `json:"quantity" xml:"quantity"`
}

func (this *Cart) Set() error {
	data, err := db.Prepare("INSERT INTO Cart(UserId, ProductId, Quantity) VALUES(?, ?, ?);")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Userid, this.ProductId, this.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (this *Cart) Get() ([]responses.CartRes, error) {
	lc := []responses.CartRes{}

	results, err := db.Query(queries.GetCart(strconv.Itoa(this.Userid)))
	if err != nil {
		return nil, err
	}

	for results.Next() {
		c := responses.CartRes{}
		err = results.Scan(&c.Id, &c.Name, &c.Quantity)
		if err != nil {
			return nil, err
		}
		lc = append(lc, c)
	}

	return lc, nil
}

func (this *Cart) Del() error {
	data, err := db.Prepare("Delete from Cart where UserId = ?")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Userid)
	if err != nil {
		return err
	}
	return nil
}

func (this *Cart) RemoveItem() error {
	data, err := db.Prepare("Delete from Cart where UserId = ? and ProductId = ?")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Userid, this.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Cart) UpdateItem() error {
	data, err := db.Prepare("Update Cart set Quantity = ? where UserId = ? and ProductId = ?")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Quantity, this.Userid, this.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Cart) CheckExist() (bool, error) {
	var check bool
	err = db.QueryRow(queries.CheckCartExist(strconv.Itoa(this.Userid), strconv.Itoa(this.ProductId))).Scan(&check)
	if err != nil {
		return false, err
	}
	return check, nil
}
