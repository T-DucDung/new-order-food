package queries

func GetListImport() string {
	return "select i.Id, i.LastUpDate, i.Total, u.Name, v.Name from Import i , Vendor v , Users u where i.IdAdmin = u.Id AND v.Id  = i.VendorId"
}
