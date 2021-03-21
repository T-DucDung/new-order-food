package requests

type RequestCreateAccount struct {
	UserName string `json:"user_name" xml:"user_name""`
	Pass     string `json:"pass" xml:"pass"`
	Type     string `json:"type" xml:"type"`
	Name     string `json:"name" xml:"name"`
	Phone    string `json:"phone" xml:"phone"`
	Email    string `json:"email" xml:"email"`
	Image    string `json:"image" xml:"image"`
	Gender   string `json:"gender" xml:"gender"`
}

type RequestUpdateStatus struct {
	IdUser string `json:"id_user" xml:"id_user"`
	Status bool   `json:"status" xml:"status"`
}
