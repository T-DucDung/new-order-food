package models

type User struct {
	Id     int    `json:"id" xml:"id"`
	Name   string `json:"name" xml:"name"`
	Phone  string `json:"phone" xml:"phone"`
	Email  string `json:"email" xml:"email"`
	Image  string `json:"image" xml:"image"`
	Gender string `json:"gender" xml:"gender"`
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

//func (this *User) GetUser()