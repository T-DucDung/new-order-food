package requests

type RequestVendor struct {
	Name    string `json:"name" xml:"name"`
	Address string `json:"address" xml:"address"`
	Email   string `json:"email" xml:"email"`
	Phone   string `json:"phone" xml:"phone"`
}
