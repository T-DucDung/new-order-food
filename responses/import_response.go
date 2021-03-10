package responses

type ImportRes struct {
	Id         int               `json:"id"`
	VendorName string            `json:"vendor_name"`
	LastUpDate int64             `json:"last_up_date"`
	NameAdmin  string            `json:"name_admin"`
	Total      float32           `json:"total"`
	Detail     []ImportDetailRes `json:"detail"`
}

type ImportDetailRes struct {
	Name     string  `json:"name"`
	Unit     string  `json:"unit"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}
