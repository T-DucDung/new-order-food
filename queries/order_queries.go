package queries

func GetListOrder(uid string) string {
	return "select o.Id, o.Name, o.Phone, o.Address, o.Total, o.CurrentStatus, o.LastUpDate from `Order` o where o.UserId = " + uid + " order by o.CurrentStatus"
}

func GetListOrderDetail(id string) string {
	return "select p.Name, od.IsSale, od.Price, od.Quantity from OrderDetail od, Product p where od.ProductId = p.Id and od.OrderId = " + id
}

func GetListOrderForAdmin() string {
	return "select o.Id, o.Name, o.Phone, o.Address, o.Total, o.CurrentStatus, o.LastUpDate from `Order` o order by o.CurrentStatus"
}
