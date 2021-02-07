package responses

type ProductRes struct {
	Id          int
	Name        string
	Image       string
	Price       float32
	IsSale      bool
	Unit        string
	Remaining   int
	SalePrice   float32
	Description string
	Sold        int
	CategoryId  int
	RateAvg     float32
}
