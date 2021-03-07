package responses

type DiscountRes struct {
	Rank       int   `json:"rank" xml:"rank"`
	Rate       int   `json:"rate" xml:"rate"`
	IdAdmin    int   `json:"id_admin" xml:"id_admin"`
	LastUpDate int64 `json:"last_up_date" xml:"last_up_date"`
}
