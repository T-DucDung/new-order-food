package models

import (
	"new-order-food/helps"
	"new-order-food/queries"
)

func CreateProduct(m map[string]string) error {
	// tao o day
	return nil
}

func GetProduct(id string) (map[string]string, error) {
	bData, err := GetDataByQuery(queries.GetProductById(id))
	if err != nil {
		return nil, err
	}
	return  helps.ByteToPrint(bData[0]), nil
}
