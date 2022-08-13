package responses

//go:generate easytags $GOFILE json,xml,bson

type ProductRes struct {
	Id          int     `json:"id" xml:"id" bson:"id"`
	Name        string  `json:"name" xml:"name" bson:"name"`
	Image       string  `json:"image" xml:"image" bson:"image"`
	Price       float64 `json:"price" xml:"price" bson:"price"`
	IsSale      bool    `json:"is_sale" xml:"is_sale" bson:"is_sale"`
	Unit        string  `json:"unit" xml:"unit" bson:"unit"`
	Remaining   int     `json:"remaining" xml:"remaining" bson:"remaining"`
	SalePrice   float64 `json:"sale_price" xml:"sale_price" bson:"sale_price"`
	Description string  `json:"description" xml:"description" bson:"description"`
	Sold        int     `json:"sold" xml:"sold" bson:"sold"`
	CategoryId  int     `json:"category_id" xml:"category_id" bson:"category_id"`
	RateAvg     float64 `json:"rate_avg" xml:"rate_avg" bson:"rate_avg"`
}
