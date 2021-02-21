package services

import (
	"new-order-food/models"
	"new-order-food/requests"
)

func CreateAccount(req requests.RequestCreateAccount) error {
	ad := models.Admin{}
	return ad.CreateAccount(req)
}
