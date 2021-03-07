package queries

func GetAllVendor() string {
	return "select v.Id, v.Name, v.Address, v.Phone, v.Email from Vendor v"
}
