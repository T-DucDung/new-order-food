package requests

//go:generate easytags $GOFILE json,xml,bson

type RequestCart struct {
	ProductId int `json:"product_id" xml:"product_id" bson:"product_id"`
	Quantity  int `json:"quantity" xml:"quantity" bson:"quantity"`
}
