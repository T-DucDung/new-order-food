package queries

func GetListUser() string {
	return "select Account.UserName, Users.Name, Users.Phone, Users.Email, Users.Image, Users.Gender from Account, Users where Account.UserId = Users.Id and Account.Type = 'user'"
}

func GetListAdmin() string {
	return "select Account.UserName, Users.Name, Users.Phone, Users.Email, Users.Image, Users.Gender from Account, Users where Account.UserId = Users.Id and Account.Type = 'admin'"
}