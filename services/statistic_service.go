package services

import (
	"new-order-food/models"
	"new-order-food/responses"
)

func TotalRevenu() (float32, error) {
	s := models.Statistics{}
	return s.TotalRevenu()
}

func TotalOrder() (int, error) {
	s := models.Statistics{}
	return s.TotalOrder()
}

func TopSale(num int) ([]responses.ProductRes, error) {
	s := models.Statistics{}
	return s.TopSale(num)
}

func TotalNewAcc() (int, error) {
	s := models.Statistics{}
	return s.TotalNewAcc()
}

func ExOrder(startTime, endTime int64) (string, error) {
	s := models.Statistics{}
	return s.ExOrder(startTime, endTime)
}

func ExImport(startTime, endTime int64) (string, error) {
	s := models.Statistics{}
	return s.ExImport(startTime, endTime)
}
