package responses

type CartRes struct {
	Name     string `json:"name" xml:"name"`
	Quantity int    `json:"quantity" xml:"quantity"`
}
