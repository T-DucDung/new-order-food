package queries

func GetAccount(username string) string {
	return "select UserName , Pass, Id, Type from Account where UserName = '" + username + "'"
}
