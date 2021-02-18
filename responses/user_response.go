package responses

type UserRes struct {
	UserName	string `json:"username" xml:"username"`
	Name		string `json:"name" xml:"name"`
	Phone		string `json:"phone" xml:"phone"`
	Image		string `json:"image" xml:"image"`
	Gender		string `json:"gender" xml:"gender"`
}
