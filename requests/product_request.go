package requests

type RequestCreateProduct struct {
	Name        string  `json:"name" xml:"name"`
	Image       string  `json:"image" xml:"image"`
	Price       float32 `json:"price" xml:"price"`
	IsSale      bool    `json:"is_sale" xml:"is_sale"`
	Unit        string  `json:"unit" xml:"unit"`
	Remaining   int     `json:"remaining" xml:"remaining"`
	SalePrice   float32 `json:"sale_price" xml:"sale_price"`
	Description string  `json:"description" xml:"description"`
	Sold        int     `json:"sold" xml:"sold"`
	CategoryId  int     `json:"category_id" xml:"category_id"`
}
