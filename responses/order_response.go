package responses

type OrderRes struct {
	Id            int              `json:"id" xml:"id"`
	Number        string           `json:"number" xml:"number"`
	Address       string           `json:"address" xml:"address"`
	Name          string           `json:"name" xml:"name"`
	LastUpdate    int64            `json:"last_update" xml:"last_update"`
	Total         float32          `json:"total" xml:"total"`
	CurrentStatus string              `json:"current_status" xml:"current_status"`
	Detail        []OrderDetailRes `json:"detail" xml:"detail"`
}

type OrderDetailRes struct {
	Name     string  `json:"name" xml:"name"`
	IsSale   bool    `json:"is_sale" xml:"is_sale"`
	Price    float32 `json:"price" xml:"price"`
	Quantity int     `json:"quantity" xml:"quantity"`
}
