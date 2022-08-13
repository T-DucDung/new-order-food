package responses

//go:generate easytags $GOFILE json,xml,bson

type CartRes struct {
	Id        int     `json:"id" xml:"id" bson:"id"`
	Name      string  `json:"name" xml:"name" bson:"name"`
	Quantity  int     `json:"quantity" xml:"quantity" bson:"quantity"`
	Price     float64 `json:"price" xml:"price" bson:"price"`
	IsSale    bool    `json:"is_sale" xml:"is_sale" bson:"is_sale"`
	SalePrice float64 `json:"sale_price" xml:"sale_price" bson:"sale_price"`
}
