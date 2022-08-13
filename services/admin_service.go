package services

import (
	"new-order-food/models"
	"new-order-food/requests"
)

func CreateAccount(req requests.RequestCreateAccount) error {
	ad := models.Admin{}
	return ad.CreateAccount(req)
}

func GetAllUser(pos, count int, status string) ([]models.User, int, error) {
	a := models.Admin{}
	if status != "" && (status == "true" || status == "false") {
		lu, err := a.GetAllUser(status)
		if err != nil {
			return nil, 0, err
		}
		if len(lu) < count {
			return lu, len(lu), nil
		}
		return lu[pos : pos+count], len(lu), nil
	}
	lu, err := a.GetAllUser("")
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return lu, 0, nil
	}
	if len(lu) < count {
		return lu, len(lu), nil
	}
	return lu[pos : pos+count], len(lu), nil
}

func UpdateStatus(req requests.RequestUpdateStatus) error {
	return (&models.Admin{}).UpdateStatus(req)
}
