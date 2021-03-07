package models

import (
	"new-order-food/queries"
)

type Vendor struct {
	Id      int    `json:"id" xml:"id"`
	Name    string `json:"name" xml:"name"`
	Address string `json:"address" xml:"address"`
	Email   string `json:"email" xml:"email"`
	Phone   string `json:"phone" xml:"phone"`
}

func (this *Vendor) GetListVendor() ([]Vendor, error) {
	lv := []Vendor{}

	results, err := db.Query(queries.GetAllVendor())
	if err != nil {
		return nil, err
	}

	for results.Next() {
		v := Vendor{}
		err = results.Scan(&v.Id, &v.Name, &v.Address, &v.Phone, &v.Email)
		if err != nil {
			return nil, err
		}
		lv = append(lv, v)
	}

	return lv, nil
}

func (this *Vendor) UpdateVendor() error {
	data, err := db.Prepare("UPDATE Vendor as v SET v.Name = ?, v.Address = ?, v.Phone = ?, v.Email = ? WHERE v.Id = ?;")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Name, this.Address, this.Phone, this.Email, this.Id)
	if err != nil {
		return err
	}

	return nil
}

func (this *Vendor) CreateVendor() error {
	data, err := db.Prepare("insert into Vendor (Name,Phone,Address,Email) VALUES(?, ?, ?, ?);")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Name, this.Phone, this.Address, this.Email)
	if err != nil {
		return err
	}
	return nil
}
