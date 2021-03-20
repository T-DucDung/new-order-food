package queries

func GetListImport() string {
	return "select i.Id, i.LastUpDate, i.Total, u.Name, v.Name from Import i , Vendor v , Users u where i.IdAdmin = u.Id AND v.Id  = i.VendorId ORDER by i.LastUpDate desc"
}

func GetListImportDetail(id string) string {
	return "select p.Name, id.Quantity, id.Price from ImportDetail id , Product p where id.ProductId = p.Id and id.IdImport = " + id
}
