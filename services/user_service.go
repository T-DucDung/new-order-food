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

func GetUser(id int) (models.User, error) {
	u := models.User{Id: id}
	return u.GetUser()
}

func GetRank(id int) (int, error) {
	u, err := GetUser(id)
	return u.Rank, err
}
