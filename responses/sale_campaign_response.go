package responses

type SaleCampaignRes struct {
	Id         int                     `json:"id" xml:"id"`
	NameAdmin  string                  `json:"name_admin" xml:"name_admin"`
	LastUpDate int64                   `json:"last_up_date" xml:"last_up_date"`
	Status     bool                    `json:"status" xml:"status"`
	Detail     []DetailSaleCampaignRes `json:"detail" xml:"detail"`
}

type DetailSaleCampaignRes struct {
	Name      string  `json:"name" xml:"name"`
	SalePrice float32 `json:"sale_price" xml:"sale_price"`
}
