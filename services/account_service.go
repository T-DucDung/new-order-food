package services

import (
	"new-order-food/models"
	"new-order-food/requests"
)

func Login(req requests.RequestLogin) (string, string, error) {
	acc := models.Account{}
	return acc.Login(req)
}

func Register(req requests.RequestRegister) error {
	acc := models.Account{}
	return acc.Register(req)
}

func CheckExistToken(token string) (int64, error) {
	return models.Exist(token)
}
