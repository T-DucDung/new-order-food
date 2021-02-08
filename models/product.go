package models

import (
	"new-order-food/queries"
	"new-order-food/responses"
)

type Product struct {
	Id          int
	Name        string
	Image       string
	Price       float32
	IsSale      bool
	Unit        string
	Remaining   int
	SalePrice   float32
	Description string
	Sold        int
	CategoryId  int
	Rate1       int
	Rate2       int
	Rate3       int
	Rate4       int
	Rate5       int
	RateAvg     float32
}

func (this *Product) CreateProduct() error {
	data, err := db.Prepare("INSERT INTO Product(Name, Image, Price, IsSale, Unit, Remaining, SalePrice, Description, Sold, CategoryId, Rate1, Rate2, Rate3, Rate4, Rate5, RateAvg) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Name, this.Image, this.Price, this.IsSale, this.Unit, this.Remaining, this.SalePrice, this.Description, this.Sold, this.CategoryId, this.Rate1, this.Rate2, this.Rate3, this.Rate4, this.Rate5, this.RateAvg)
	if err != nil {
		return err
	}
	return nil
}

func (this *Product) GetProduct(id string) (responses.ProductRes, error) {
	p := responses.ProductRes{}
	err = db.QueryRow(queries.GetProductById(id)).Scan(&p.Id, &p.Name, &p.Image, &p.Price, &p.IsSale, &p.Unit, &p.Remaining, &p.SalePrice, &p.Description, &p.Sold, &p.CategoryId, &p.RateAvg)
	if err != nil {
		return responses.ProductRes{}, err
	}
	return p, nil
}

func (this *Product) GetListProduct(query string) ([]responses.ProductRes, error) {
	lp := []responses.ProductRes{}

	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		p := responses.ProductRes{}
		err = results.Scan(&p.Id, &p.Name, &p.Image, &p.Price, &p.IsSale, &p.Unit, &p.Remaining, &p.SalePrice, &p.Description, &p.Sold, &p.CategoryId, &p.RateAvg)
		if err != nil {
			return nil, err
		}
		lp = append(lp, p)
	}

	return lp, nil
}

func (this *Product) UpDateProduct() error {
	data, err := db.Prepare("UPDATE Product as p SET p.Name = ?, p.Image = ?, p.Price = ?, p.IsSale = ?, p.Unit = ?, p.Remaining = ?, p.SalePrice = ?, p.Description = ?, p.Sold = ?, p.CategoryId = ?, p.Rate1 = ?, p.Rate2 = ?, p.Rate3 = ?, p.Rate4 = ?, p.Rate5 = ?, p.RateAvg = ? WHERE p.Id = ?;")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Name, this.Image, this.Price, this.IsSale, this.Unit, this.Remaining, this.SalePrice, this.Description, this.Sold, this.CategoryId, this.Rate1, this.Rate2, this.Rate3, this.Rate4, this.Rate5, this.RateAvg, this.Id)
	if err != nil {
		return err
	}

	return nil
}
