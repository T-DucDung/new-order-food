package services

import (
	"new-order-food/models"
	"new-order-food/requests"
	"new-order-food/responses"
)

func CreateSale(req requests.RequestSale, id int) error {
	s := models.SaleCampaign{
		IdAdmin: id,
		Status:  true,
	}
	sd := []models.DetailSaleCampaign{}
	for _, item := range req.Detail {
		s := models.DetailSaleCampaign{
			ProductId: item.ProductId,
			SalePrice: item.SalePrice,
		}
		sd = append(sd, s)
	}
	s.Detail = sd

	return s.Create()
}

func GetSale() ([]responses.SaleCampaignRes, error) {
	s := models.SaleCampaign{}
	return s.GetList()
}

func UpdateSale(id int) error {
	s := models.SaleCampaign{
		Id:     id,
		Status: false,
	}
	return s.UpDateStatus()
}
