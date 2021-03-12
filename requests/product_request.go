package requests

type RequestCreateProduct struct {
	Name        string  `json:"name" xml:"name"`
	Image       string  `json:"image" xml:"image"`
	Unit        string  `json:"unit" xml:"unit"`
	Description string  `json:"description" xml:"description"`
	CategoryId  int     `json:"category_id" xml:"category_id"`
}
