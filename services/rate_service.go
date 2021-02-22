package services

import (
	"new-order-food/models"
)

func SetRate(uid, pid int, rate int) error {
	r := models.Rate{}
	return r.SetRate(pid, uid, rate)
}
