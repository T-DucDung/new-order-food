package requests

type RequestSale struct {
	Detail     []RequestDetailSale `json:"detail" xml:"detail"`
}

type RequestDetailSale struct {
	ProductId int     `json:"product_id" xml:"product_id"`
	SalePrice float32 `json:"sale_price" xml:"sale_price"`
}
