package queries

func GetListCate() string {
	return "select c.Name, COUNT(p.Id) as total FROM Category c left join Product p on c.CategoryId = p.CategoryId GROUP BY c.CategoryId , c.Name"
}
