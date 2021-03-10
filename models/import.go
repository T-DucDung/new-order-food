package models

type Import struct {
	Id         int            `json:"id" xml:"id"`
	VendorId   int            `json:"vendor_id" xml:"vendor_id"`
	LastUpDate int64          `json:"last_up_date" xml:"last_up_date"`
	IdAdmin    int            `json:"id_admin" xml:"id_admin"`
	Total      float32        `json:"total" xml:"total"`
	Detail     []ImportDetail `json:"detail" xml:"detail"`
}

type ImportDetail struct {
	ProductId int     `json:"product_id" xml:"product_id"`
	Unit      string  `json:"unit" xml:"unit"`
	Quantity  int     `json:"quantity" xml:"quantity"`
	Price     float32 `json:"price" xml:"price"`
}

