package services

import (
	"new-order-food/models"
	"new-order-food/requests"
	"new-order-food/responses"
)

func Import(req requests.RequestImport, id int) error {
	i := models.Import{
		Id:       0,
		VendorId: req.VendorId,
		IdAdmin:  id,
		Total:    0,
		Detail:   nil,
	}

	newDetail := []models.ImportDetail{}
	var total float32
	total = 0

	for _, item := range req.Detail {
		detail := models.ImportDetail{
			ProductId: item.ProductId,
			Unit:      item.Unit,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
		total = total + (item.Price * float32(item.Quantity))
		newDetail = append(newDetail, detail)
	}
	i.Total = total
	i.Detail = newDetail

	return i.Import()
}

func GetAllImport() ([]responses.ImportRes, error) {
	i := models.Import{}
	return i.GetAllImport()
}
