package services

import (
	"new-order-food/models"
	"new-order-food/requests"
	"new-order-food/responses"
	"strconv"
	"time"
)

func GetListDiscount() ([]responses.DiscountRes, error) {
	c := models.Discount{}
	return c.GetListDiscount()
}

func UpdateDiscount(req requests.RequestDis, uid int) error {
	discount := models.Discount{
		Rank:       req.Rank,
		Rate:       req.Rate,
		IdAdmin:    uid,
		LastUpDate: time.Now().Unix(),
	}
	return discount.UpDateDiscount()
}

func CreateDiscount(req requests.RequestDis, uid int) error {
	discount := models.Discount{
		Rank:       req.Rank,
		Rate:       req.Rate,
		IdAdmin:    uid,
		LastUpDate: time.Now().Unix(),
	}
	return discount.CreateDiscount()
}

func GetRateDis(rank int) (float32, error) {
	dis := models.Discount{}
	return dis.GetRateDis(strconv.Itoa(rank))
}
