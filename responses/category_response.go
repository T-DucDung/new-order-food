package responses

type CategoryRes struct {
	Name  string `json:"name" xml:"name"`
	Total int    `json:"total" xml:"total"`
}
