package requests

type RequestImport struct {
	VendorId   int                   `json:"vendor_id" xml:"vendor_id"`
	Total      float32               `json:"total" xml:"total"`
	Detail     []RequestImportDetail `json:"detail" xml:"detail"`
}

type RequestImportDetail struct {
	ProductId int     `json:"product_id" xml:"product_id"`
	Quantity  int     `json:"quantity" xml:"quantity"`
	Price     float32 `json:"price" xml:"price"`
}
