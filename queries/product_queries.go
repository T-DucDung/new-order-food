package queries

func GetProductById(id string) string {
	return "select p.Id, p.Name, p.Image, p.Price, case when p.IsSale = 1 then 'true' else 'false' end as IsSale, p.Unit, p.Remaining, p.SalePrice, p.Description, p.Sold, p.CategoryId, p.RateAvg from Product p where id = " + id
}