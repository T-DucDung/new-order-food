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

func GetListProduct(pos, count int, cateid string) ([]responses.ProductRes, error) {
	p := models.Product{}
	if cateid != "" {
		return p.GetListProduct(queries.GetListProductByCate(cateid))
	}
	lp, err := p.GetListProduct(queries.GetListProduct())
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return lp, nil
	}
	return lp[pos : count+1], nil
}

func UpDateProduct(p models.Product) error {
	return p.UpDateProduct()
}
