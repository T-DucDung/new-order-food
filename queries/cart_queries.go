package queries

func GetCart(userid string) string {
	return "select p.Id, p.Name, c.Quantity from Product p , Cart c where p.Id = c.ProductId and c.UserId = '" + userid + "'"
}

func CheckExist(uid, pid string) string {
	return "select exists(select ProductId from Cart where UserId = " + uid + " and ProductId = " + pid + ")"
}
