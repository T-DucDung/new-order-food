package requests

//go:generate easytags $GOFILE json,xml,bson

type RequestCreateAccount struct {
	UserName string `json:"user_name" xml:"user_name" bson:"user_name"`
	Pass     string `json:"pass" xml:"pass" bson:"pass"`
	Type     string `json:"type" xml:"type" bson:"type"`
	Name     string `json:"name" xml:"name" bson:"name"`
	Phone    string `json:"phone" xml:"phone" bson:"phone"`
	Email    string `json:"email" xml:"email" bson:"email"`
	Image    string `json:"image" xml:"image" bson:"image"`
	Gender   string `json:"gender" xml:"gender" bson:"gender"`
}

type RequestUpdateStatus struct {
	IdUser int  `json:"id_user" xml:"id_user" bson:"id_user"`
	Status bool `json:"status" xml:"status" bson:"status"`
}
