package services

import (
	"new-order-food/models"
	"new-order-food/queries"
	"new-order-food/requests"
	"new-order-food/responses"
)

func CreateProduct(req requests.RequestCreateProduct) error {
	p := models.Product{
		Name:        req.Name,
		Image:       req.Image,
		Price:       req.Price,
		IsSale:      req.IsSale,
		Unit:        req.Unit,
		Remaining:   req.Remaining,
		SalePrice:   req.SalePrice,
		Description: req.Description,
		Sold:        req.Sold,
		CategoryId:  req.CategoryId,
		Rate1:       0,
		Rate2:       0,
		Rate3:       0,
		Rate4:       0,
		Rate5:       0,
		RateAvg:     0,
	}
	return p.CreateProduct()
}

func GetProduct(id string) (responses.ProductRes, error) {
	p := models.Product{}
	return p.GetProduct(id)
}

func GetListProduct(cateid string) ([]responses.ProductRes, error) {
	p := models.Product{}
	if cateid != "" {
		return p.GetListProduct(queries.GetListProductByCate(cateid))
	}
	return p.GetListProduct(queries.GetListProduct())
}

func UpDateProduct(p models.Product) error {
	return p.UpDateProduct()
}
