package queries

func GetListCate() string {
	return "select c.Name, COUNT(p.Id) as total FROM Category c, Product p where c.CategoryId = p.CategoryId GROUP BY c.CategoryId , c.Name"
}
