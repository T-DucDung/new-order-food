package queries

import "fmt"

func GetAllDiscount() string {
	return "select d.`Rank`, d.Rate, d.IdAdmin, d.LastUpDate, d.Accumulate from Discount d"
}

func GetRateDis(rank string) string {
	return "select d.Rate from Discount d where d.`Rank` = " + rank
}

func GetRank(total float32) string {
	return "select d.`Rank` from Discount d where Accumulate < " + fmt.Sprintf("%f", total) + " order by `Rank` desc limit 1"
}
