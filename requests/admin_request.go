package requests

type RequestCreateAccount struct {
	Username string `json:"username" xml:"username"`
	Pass     string `json:"pass" xml:"pass"`
	Type     string `json:"type" xml:"type"`
	Name     string `json:"name" xml:"name"`
	Phone    string `json:"phone" xml:"phone"`
	Email    string `json:"email" xml:"email"`
	Image    string `json:"image" xml:"image"`
	Gender   string `json:"gender" xml:"gender"`
}
