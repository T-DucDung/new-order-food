package models

import (
	"new-order-food/queries"
	"new-order-food/responses"
	"strconv"
)

type Comment struct {
	Id         int    `json:"id" xml:"id"`
	UserId     int    `json:"user_id" xml:"user_id"`
	ProductId  int    `json:"product_id" xml:"product_id"`
	Comment    string `json:"comment" xml:"comment"`
	LastUpDate int64  `json:"last_up_date" xml:"last_up_date"`
}

func (this *Comment) CreateComment() error {
	data, err := db.Prepare("insert into Comment (UserId,ProductId,Comment,LastUpDate) VALUES(?, ?, ?, ?);")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.UserId, this.ProductId, this.Comment, this.LastUpDate)
	if err != nil {
		return err
	}
	return nil
}

func (this *Comment) Get() ([]responses.CommentRes, error) {
	lc := []responses.CommentRes{}

	results, err := db.Query(queries.GetComment(strconv.Itoa(this.ProductId)))
	if err != nil {
		return nil, err
	}

	for results.Next() {
		c := responses.CommentRes{}
		err = results.Scan(&c.UserId, &c.Name, &c.Comment, &c.LastUpDate)
		if err != nil {
			return nil, err
		}
		lc = append(lc, c)
	}

	return lc, nil
}

func (this *Comment) RemoveItem() error {
	data, err := db.Prepare("Delete from Comment where UserId = ? and ProductId = ?")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.UserId, this.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Comment) UpdateItem() error {
	data, err := db.Prepare("Update Comment set Comment = ? where UserId = ? and ProductId = ?")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Comment, this.UserId, this.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (this *Comment) CheckExist() (bool, error) {
	var check bool
	err = db.QueryRow(queries.CheckCommentExist(strconv.Itoa(this.UserId), strconv.Itoa(this.ProductId))).Scan(&check)
	if err != nil {
		return false, err
	}
	return check, nil
}
