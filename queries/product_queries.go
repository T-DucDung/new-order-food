package queries

func GetProductById(id string) string {
	return "select p.Id, p.Name, p.Image, p.Price, p.IsSale, p.Unit, p.Remaining, p.SalePrice, p.Description, p.Sold, p.CategoryId, p.RateAvg from Product p where id = " + id
}

func GetListProduct() string {
	return "select p.Id, p.Name, p.Image, p.Price, p.IsSale, p.Unit, p.Remaining, p.SalePrice, p.Description, p.Sold, p.CategoryId, p.RateAvg from Product p"
}