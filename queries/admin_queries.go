package queries

func GetALlUser() string {
	return "select u.Id, u.Name, u.Gender, u.Image, u.Email, u.Phone, u.`Rank` from Users u"
}
