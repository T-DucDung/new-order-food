package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"new-order-food/models"
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

func GetProduct(id int) (responses.ProductRes, error) {
	return (&models.Product{}).GetProduct(id)
}

func GetListProduct(pos, count int, cateid int) ([]responses.ProductRes, int, error) {
	p := models.Product{}
	if cateid != -1 {
		lp, err := p.GetListProduct(bson.D{{"category_id", cateid}})
		if err != nil {
			return nil, 0, err
		}
		if len(lp) < count {
			return lp, len(lp), nil
		}
		return lp[pos : pos+count], len(lp), nil
	}
	lp, err := p.GetListProduct(bson.D{})
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

func UpdateProduct(p models.Product) error {
	return p.UpDateProduct()
}

