package services

import (
	"new-order-food/models"
	"new-order-food/requests"
)

func CreateProduct(req requests.RequestCreateProduct) error {

 	return nil
}

func GetProduct(id string) (map[string]string, error) {
	p := models.Product{}
	return p.GetProduct(id)
}
