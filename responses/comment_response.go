package responses

type CommentRes struct {
	UserId     int    `json:"user_id" xml:"user_id"`
	Name       string `json:"name" xml:"name"`
	Comment    string `json:"comment" xml:"comment"`
	LastUpDate int64  `json:"last_up_date" xml:"last_up_date"`
}
