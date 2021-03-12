package queries

func GetProductById(id string) string {
	return "select p.Id, p.Name, p.Image, p.Price, p.IsSale, p.Unit, p.Remaining, p.SalePrice, p.Description, p.Sold, p.CategoryId, p.RateAvg from Product p where id = " + id
}

func GetListProduct() string {
	return "select p.Id, p.Name, p.Image, p.Price, p.IsSale, p.Unit, p.Remaining, p.SalePrice, p.Description, p.Sold, p.CategoryId, p.RateAvg from Product p"
}

func GetListProductByCate(cate string) string {
	return "select p.Id, p.Name, p.Image, p.Price, p.IsSale, p.Unit, p.Remaining, p.SalePrice, p.Description, p.Sold, p.CategoryId, p.RateAvg from Product p where p.CategoryId = " + cate
}

func GetTotalRemaining(pid string) string {
	return "select p.Remaining FROM Product p where p.Id = " + pid
}

func GetPrice(pid string) string {
	return "select p.Price, p.IsSale, p.SalePrice from Product p where p.Id = " + pid
}

func GetRateProduct(pid string) string {
	return "select p.RateAvg FROM Product p where p.Id = " + pid
}

func GetAllRate(pid string) string {
	return "select p.Rate1, p.Rate2, p.Rate3, p.Rate4, p.Rate5 FROM Product p where p.Id = " + pid
}

func GetExist(pid string) string {
	return ("select exists(select p.Id from Product p where p.Id = " + pid + ")")
}
