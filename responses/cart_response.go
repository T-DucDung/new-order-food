package responses

type CartRes struct {
	Id        int     `json:"id" xml:"id"`
	Name      string  `json:"name" xml:"name"`
	Quantity  int     `json:"quantity" xml:"quantity"`
	Price     float32 `json:"price" xml:"price"`
	IsSale    bool    `json:"is_sale" xml:"is_sale"`
	SalePrice float32 `json:"sale_price" xml:"sale_price"`
}
