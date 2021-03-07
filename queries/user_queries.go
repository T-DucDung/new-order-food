package queries

func GetInfoUser(id string) string {
	return "select u.Id, u.Name, u.Gender, u.Image, u.Email, u.Phone, u.`Rank` from Users u where u.Id = " + id
}

func GetAllUser() string {
	return "select u.Id, u.Name, u.Gender, u.Image, u.Email, u.Phone, u.`Rank` from Users u"
}

func GetAllUserByStatus(status string) string {
	return "select u.Id, u.Name, u.Gender, u.Image, u.Email, u.Phone, u.`Rank` from Users u, Account a where u.Id = a.Id and a.Status = " + status
}
