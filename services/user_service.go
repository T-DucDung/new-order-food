package services

import (
	"new-order-food/models"
	"new-order-food/requests"
)

func UpdateUser(req requests.RequestUser, uid int) error {
	u := models.User{
		Id:     uid,
		Name:   req.Name,
		Phone:  req.Phone,
		Email:  req.Email,
		Image:  req.Image,
		Gender: req.Gender,
	}
	return u.UpdateUser()
}
