package models

import (
	"new-order-food/queries"
	"new-order-food/responses"
)

type Category struct {
	CategoryId int    `json:"category_id" xml:"category_id"`
	Name       string `json:"name" xml:"name"`
}

func (this *Category) GetListCategory() ([]responses.CategoryRes, error) {
	lc := []responses.CategoryRes{}

	results, err := db.Query(queries.GetListCate())
	if err != nil {
		return nil, err
	}

	for results.Next() {
		p := responses.CategoryRes{}
		err = results.Scan(&p.Name, &p.Total)
		if err != nil {
			return nil, err
		}
		lc = append(lc, p)
	}

	return lc, nil
}

func (this *Category) UpDateCategory() error {
	data, err := db.Prepare("UPDATE Category as c SET c.Name = ? WHERE c.CategoryId = ?;")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Name, this.CategoryId)
	if err != nil {
		return err
	}

	return nil
}

func (this *Category) CreateCategory() error {
	data, err := db.Prepare("INSERT INTO Category(Name) VALUES(?);")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Name)
	if err != nil {
		return err
	}
	return nil
}
