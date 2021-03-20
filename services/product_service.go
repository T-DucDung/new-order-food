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
		Price:       0,
		IsSale:      false,
		Unit:        req.Unit,
		Remaining:   0,
		SalePrice:   0,
		Description: req.Description,
		Sold:        0,
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

func GetListProduct(pos, count int, cateid string) ([]responses.ProductRes, int, error) {
	p := models.Product{}
	if cateid != "" {
		lp, err := p.GetListProduct(queries.GetListProductByCate(cateid))
		if err != nil {
			return nil, 0, err
		}
		if len(lp) < count {
			return lp, len(lp), nil
		}
		return lp[pos : pos+count], len(lp), nil
	}
	lp, err := p.GetListProduct(queries.GetListProduct())
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return lp, 0, nil
	}
	if len(lp) < count {
		return lp, len(lp), nil
	}
	return lp[pos : pos+count], len(lp), nil
}

func UpDateProduct(p models.Product) error {
	return p.UpDateProduct()
}

func Search(word string) ([]responses.ProductRes, error) {
	p := models.Product{}
	return p.Search(word)
}
