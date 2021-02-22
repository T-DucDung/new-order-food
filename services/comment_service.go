package services

import (
	"new-order-food/models"
	"new-order-food/requests"
	"new-order-food/responses"
	"time"
)

func CreateComment(req requests.RequestComment, uid int) error {
	com := models.Comment{
		UserId:    uid,
		ProductId: req.ProductId,
		Comment:   req.Comment,
		LastUpDate: time.Now().Unix(),
	}
	return com.CreateComment()
}

func GetComments(pid int) ([]responses.CommentRes, error) {
	com := models.Comment{
		ProductId: pid,
	}
	return com.Get()
}
