package models

import (
	"new-order-food/helps"
	"new-order-food/queries"
)

type Product struct {
	Id          int
	Name        string
	Image       string
	Price       float32
	IsSale      bool
	Unit        string
	Remaining   int
	SalePrice   float32
	Description string
	Sole        int
	CategoryId  int
	Rate1       int
	Rate2       int
	Rate3       int
	Rate4       int
	Rate5       int
	RateAvg     float32
}

func (this *Product) CreateProduct(m map[string]string) error {
	// tao o day
	return nil
}

func (this *Product) GetProduct(id string) (map[string]string, error) {
	bData, err := GetDataByQuery(queries.GetProductById(id))
	if err != nil {
		return nil, err
	}
	return helps.ByteToPrint(bData[0]), nil
}
