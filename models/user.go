package models

import (
	// "new-order-food/queries"
	"new-order-food/responses"
)

type User struct {
	Id     int    `json:"id" xml:"id"`
	Name   string `json:"name" xml:"name"`
	Phone  string `json:"phone" xml:"phone"`
	Email  string `json:"email" xml:"email"`
	Image  string `json:"image" xml:"image"`
	Gender string `json:"gender" xml:"gender"`
}


func (this *User) GetListUser(query string) ([]responses.UserRes, error) {
	lp := []responses.UserRes{}

	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		p := responses.UserRes{}
		err = results.Scan(&p.UserName, &p.Name, &p.Phone, &p.Image, &p.Gender)
		if err != nil {
			return nil, err
		}
		lp = append(lp, p)
	}

	return lp, nil
}

func (this *User) GetListAdmin(query string) ([]responses.UserRes, error) {
	lp := []responses.UserRes{}

	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		p := responses.UserRes{}
		err = results.Scan(&p.UserName, &p.Name, &p.Phone, &p.Image, &p.Gender)
		if err != nil {
			return nil, err
		}
		lp = append(lp, p)
	}

	return lp, nil
}

