package requests

type RequestComment struct {
	ProductId int    `json:"product_id" xml:"product_id"`
	Comment   string `json:"comment" xml:"comment"`
}
