package queries

func GetRate(uid, pid string) string {
	return "select r.Rate from Rate r where UserId = " + uid + " and ProductId = " + pid
}

func CheckRateExist(uid, pid string) string {
	return "select exists(select UserId from Rate where UserId = " + uid + " and ProductId = " + pid + ")"
}

func CountRate(pid string) string {
	return "select count(Rate) as count from Rate r where ProductId = " + pid
}
