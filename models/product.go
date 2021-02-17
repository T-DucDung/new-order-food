package models

import (
	"new-order-food/queries"
	"new-order-food/responses"
)

type Product struct {
	Id          int     `json:"id" xml:"id"`
	Name        string  `json:"name" xml:"name"`
	Image       string  `json:"image" xml:"image"`
	Price       float32 `json:"price" xml:"price"`
	IsSale      bool    `json:"is_sale" xml:"is_sale"`
	Unit        string  `json:"unit" xml:"unit"`
	Remaining   int     `json:"remaining" xml:"remaining"`
	SalePrice   float32 `json:"sale_price" xml:"sale_price"`
	Description string  `json:"description" xml:"description"`
	Sold        int     `json:"sold" xml:"sold"`
	CategoryId  int     `json:"category_id" xml:"category_id"`
	Rate1       int     `json:"rate1" xml:"rate1"`
	Rate2       int     `json:"rate2" xml:"rate2"`
	Rate3       int     `json:"rate3" xml:"rate3"`
	Rate4       int     `json:"rate4" xml:"rate4"`
	Rate5       int     `json:"rate5" xml:"rate5"`
	RateAvg     float32 `json:"rate_avg" xml:"rate_avg"`
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

func (this *Product) CheckRemaining(pid string) (int, error) {
	var total int
	err := db.QueryRow(queries.GetTotalRemaining(pid)).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (this *Product) UpdateRemaining(pid string, total int) error {
	data, err := db.Prepare("UPDATE Product as p SET p.Remaining = ? WHERE p.Id = ?;")
	if err != nil {
		return err
	}
	_, err = data.Exec(total, pid)
	if err != nil {
		return err
	}
	return nil
}
