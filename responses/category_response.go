package responses

type CategoryRes struct {
	Id    int `json:"id" xml:"id"`
	Name  string `json:"name" xml:"name"`
	Total int    `json:"total" xml:"total"`
}
