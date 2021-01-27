package queries

func GetProductById(id string) string {
	return "select * from product where id = " + id
}