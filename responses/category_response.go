package responses

type CategoryRes struct {
	Id    string `json:"id" xml:"id"`
	Name  string `json:"name" xml:"name"`
	Total int    `json:"total" xml:"total"`
}
