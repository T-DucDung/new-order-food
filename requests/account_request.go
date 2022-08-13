package requests

//go:generate easytags $GOFILE json,xml,bson

type RequestLogin struct {
	Username string `json:"username" xml:"username" bson:"username"`
	Pass     string `json:"pass" xml:"pass" bson:"pass"`
}

type RequestRegister struct {
	Username string `json:"username" xml:"username" bson:"username"`
	Pass     string `json:"pass" xml:"pass" bson:"pass"`
	Name     string `json:"name" xml:"name" bson:"name"`
	Phone    string `json:"phone" xml:"phone" bson:"phone"`
	Email    string `json:"email" xml:"email" bson:"email"`
	Image    string `json:"image" xml:"image" bson:"image"`
	Gender   string `json:"gender" xml:"gender" bson:"gender"`
}
