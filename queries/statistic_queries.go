package queries

import (
	"strconv"
	"time"
)

func GetTotalRevenu() string {
	start := time.Now().Truncate(time.Hour * 24).Unix()
	end := time.Now().Unix()
	return "select SUM(o.Total) from `Order` o where o.CurrentStatus = 2 and LastUpDate BETWEEN " + strconv.FormatInt(start, 10) + " and " + strconv.FormatInt(end, 10)
}

func GetTotalOrder() string {
	start := time.Now().Truncate(time.Hour * 24).Unix()
	end := time.Now().Unix()
	return "select COUNT(o.Id) from `Order` o where LastUpDate BETWEEN " + strconv.FormatInt(start, 10) + " and " + strconv.FormatInt(end, 10)
}

func GetTopSale(num int) string {
	return "SELECT p.Id, p.Name, p.Image, p.Price, p.IsSale, p.Unit, p.Remaining, p.SalePrice, p.Description, p.Sold, p.CategoryId, p.RateAvg from Product p where p.Id in (SELECT od.ProductId from OrderDetail od GROUP by od.ProductId Order by sum(od.Quantity)) LIMIT " + strconv.Itoa(num)
}

func GetTotalAcc() string {
	return "SELECT count(a.Id) from Account a where a.`Type` = 'user'"
}

func ExOrder(startTime, endTime int64) string {
	return "select o.Id, o.Name, o.Phone, o.Address, o.Total, o.CurrentStatus, o.LastUpDate from `Order` o where LastUpDate BETWEEN " + strconv.FormatInt(startTime, 10) + " and " + strconv.FormatInt(endTime, 10)
}
