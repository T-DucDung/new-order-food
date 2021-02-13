package models

type User struct {
	Id     int    `json:"id" xml:"id"`
	Name   string `json:"name" xml:"name"`
	Phone  string `json:"phone" xml:"phone"`
	Email  string `json:"email" xml:"email"`
	Image  string `json:"image" xml:"image"`
	Gender string `json:"gender" xml:"gender"`
}
