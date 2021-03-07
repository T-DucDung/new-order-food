package queries

func GetAllDiscount() string {
	return "select d.`Rank`, d.Rate, d.IdAdmin, d.LastUpDate from Discount d"
}

func GetRateDis(rank string) string {
	return "select d.Rate from Discount d where d.`Rank` = " + rank
}
