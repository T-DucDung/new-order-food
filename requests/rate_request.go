package requests

type RequestRate struct {
	ProductId int `json:"product_id" xml:"product_id"`
	Rate      int `json:"rate" xml:"rate"`
}
