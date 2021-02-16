package requests

type RequestLogin struct {
	Username string
	Pass     string
}

type RequestRegister struct {
	Username string
	Pass     string
	Name     string
	Phone    string
	Email    string
	Image    string
	Gender   string
}
