package queries

func GetAllSale() string {
	return "SELECT sc.Id, u.Name, sc.LastUpDate, sc.Status FROM SaleCampaign sc, Users u where u.Id = sc.IdAdmin order by sc.LastUpDate desc"
}

func GetDetailSale(id string) string {
	return "select p.Name , scd.SalePrice from SaleCampaignDetail scd, Product p where scd.ProductId =p.Id and scd.CampaignId = " + id
}

func GetIdProduct(id string) string {
	return "select scd.ProductId from SaleCampaignDetail scd where scd.CampaignId = " + id
}
