package models

import (
	"new-order-food/queries"
	"strconv"
)

type User struct {
	Id     int    `json:"id" xml:"id"`
	Name   string `json:"name" xml:"name"`
	Phone  string `json:"phone" xml:"phone"`
	Email  string `json:"email" xml:"email"`
	Image  string `json:"image" xml:"image"`
	Gender string `json:"gender" xml:"gender"`
	Rank   int    `json:"rank" xml:"rank"`
}

func (this *User) UpdateUser() error {
	data, err := db.Prepare("UPDATE Users as u set u.Name = ? , u.Phone = ? , u.Gender = ? , u.Email = ? , u.Image = ? WHERE u.Id = ?")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Name, this.Phone, this.Gender, this.Email, this.Image, this.Id)
	if err != nil {
		return err
	}
	return nil
}

func (this *User) GetUser() (User, error) {
	u := User{}
	err = db.QueryRow(queries.GetInfoUser(strconv.Itoa(this.Id))).Scan(&u.Id, &u.Name, &u.Gender, &u.Image, &u.Email, &u.Phone, &u.Rank)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (this *User) UpRank() error {
	data, err := db.Prepare("UPDATE Users as u set u.Rank = ? WHERE u.Id = ?")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Rank, this.Id)
	if err != nil {
		return err
	}
	return nil
}
