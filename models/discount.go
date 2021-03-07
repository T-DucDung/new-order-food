package models

import (
	"new-order-food/queries"
	"new-order-food/responses"
)

type Discount struct {
	Id         int   `json:"id" xml:"id"`
	Rank       int   `json:"rank" xml:"rank"`
	Rate       float32   `json:"rate" xml:"rate"`
	IdAdmin    int   `json:"id_admin" xml:"id_admin"`
	LastUpDate int64 `json:"last_up_date" xml:"last_up_date"`
}

func (this *Discount) GetListDiscount() ([]responses.DiscountRes, error) {
	ld := []responses.DiscountRes{}

	results, err := db.Query(queries.GetAllDiscount())
	if err != nil {
		return nil, err
	}

	for results.Next() {
		d := responses.DiscountRes{}
		err = results.Scan(&d.Rank, &d.Rate, &d.IdAdmin, &d.LastUpDate)
		if err != nil {
			return nil, err
		}
		ld = append(ld, d)
	}

	return ld, nil
}

func (this *Discount) UpDateDiscount() error {
	data, err := db.Prepare("UPDATE Discount as d SET d.Rate = ? WHERE d.Rank = ?;")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Rate, this.Rank)
	if err != nil {
		return err
	}

	return nil
}

func (this *Discount) CreateDiscount() error {
	data, err := db.Prepare("insert into Discount (`Rank` ,Rate,IdAdmin,LastUpDate) VALUES(?, ?, ?, ?);")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Rank, this.Rate, this.IdAdmin, this.LastUpDate)
	if err != nil {
		return err
	}
	return nil
}

func (this *Discount) GetRateDis(rank string) (float32, error) {
	var rate float32
	err = db.QueryRow(queries.GetRateDis(rank)).Scan(&rate)
	if err != nil {
		return 0, err
	}
	return rate, nil
}
