package requests

type RequestOrder struct {
	lastUpdate int64
	Number     string               `json:"number" xml:"number"`
	Address    string               `json:"address" xml:"address"`
	Name       string               `json:"name" xml:"name"`
	Detail     []RequestOrderDetail `json:"detail" xml:"detail"`
}

type RequestOrderDetail struct {
	ProductId int `json:"product_id" xml:"product_id"`
	Quantity  int `json:"quantity" xml:"quantity"`
}
