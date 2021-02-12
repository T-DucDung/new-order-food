package services

import (
	"new-order-food/models"
	"new-order-food/requests"
)

func Login(req requests.RequestLogin) (string, error) {
	acc := models.Account{}
	return acc.Login(req)
}
func Register(req requests.RequestRegister) error {
	acc := models.Account{}
	return acc.Register(req)
}

func RegisterForAdmin(req requests.RequestRegisterForAdmin) error {
	acc := models.Account{}
	return acc.RegisterForAdmin(req)
}
