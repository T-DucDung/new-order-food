package queries

func GetCart(userid string) string {
	return "select p.Name, c.Quantity from Product p , Cart c where p.Id = c.ProductId and c.UserId = '" + userid + "'"
}
