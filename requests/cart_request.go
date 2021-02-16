package requests

type RequestCart struct {
	ProductId int `json:"product_id" xml:"product_id"`
	Quantity  int `json:"quantity" xml:"quantity"`
}
