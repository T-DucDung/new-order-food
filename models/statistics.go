package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"new-order-food/responses"
	"sort"
	"time"
)

type Statistics struct{}

func (this *Statistics) TotalRevenu() (float32, error) {
	//var total float32
	//err = db.QueryRow(queries.GetTotalRevenu()).Scan(&total)
	//if err != nil {
	//	return -1, err
	//}
	//return total, nil
	return 0, nil
}

func (this *Statistics) TotalOrder() (int, error) {
	//var total int
	//err = db.QueryRow(queries.GetTotalOrder()).Scan(&total)
	//if err != nil {
	//	return -1, err
	//}
	//return total, nil
	return 0, nil
}

func (this *Statistics) TopSale(num int) ([]responses.ProductRes, error) {
	res := make([]responses.ProductRes, 0)
	uPid := make(map[int]int)

	cur, err := getOrderDetailDocument().Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err.Error(), "err.Error() models/statistics.go:38")
		return []responses.ProductRes{}, err
	}

	for cur.Next(context.TODO()) {
		var elem InsOrderDetail
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/statistics.go:47")
			return []responses.ProductRes{}, err
		}

		if val, ok := uPid[elem.ProductId]; ok {
			uPid[elem.ProductId] = val + elem.Quantity
		} else {
			uPid[elem.ProductId] = elem.Quantity
		}
	}

	if err := cur.Err(); err != nil {
		log.Println(err.Error(), "err.Error() models/statistics.go:61")
		return []responses.ProductRes{}, err
	}

	keys := make([]int, 0, len(uPid))
	for key := range uPid {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return uPid[keys[i]] > uPid[keys[j]] })

	for key, _ := range uPid {
		p, _ := (&Product{}).GetProduct(key)
		res = append(res, p)
		num--
		if num == 0 {
			break
		}
	}

	return res, nil
}

func (this *Statistics) TotalNewAcc() (int, error) {
	//var total int
	//err = db.QueryRow(queries.GetTotalAcc()).Scan(&total)
	//if err != nil {
	//	return -1, err
	//}
	//return total, nil

	return 0, nil
}

func (this *Statistics) ExOrder(startTime, endTime int64) (string, error) {
	//lo := []responses.OrderRes{}
	//
	//results, err := db.Query(queries.ExOrder(startTime, endTime))
	//if err != nil {
	//	return "", err
	//}
	//
	//for results.Next() {
	//	o := responses.OrderRes{}
	//	status := 0
	//	err = results.Scan(&o.Id, &o.Name, &o.Number, &o.Address, &o.Total, &status, &o.LastUpdate)
	//	if err != nil {
	//		return "", err
	//	}
	//	o.CurrentStatus = mapStatus[status]
	//	lo = append(lo, o)
	//}
	//
	//for index, item := range lo {
	//	lod, err := getOrderDetail(strconv.Itoa(item.Id))
	//	if err != nil {
	//		return "", err
	//	}
	//	lo[index].Detail = lod
	//}
	//f := excelize.NewFile()
	//sytes, _ := f.NewStyle(`{"alignment":{"vertical":"top"}}`)
	//f.SetCellStyle("Sheet1", "G2", "G"+strconv.Itoa(len(lo)+1), sytes)
	//
	//f.SetCellValue("Sheet1", "A1", "Order Id")
	//f.SetCellValue("Sheet1", "B1", "Cus Name")
	//f.SetCellValue("Sheet1", "C1", "Address")
	//f.SetCellValue("Sheet1", "D1", "Time Order")
	//f.SetCellValue("Sheet1", "E1", "Order's State")
	//f.SetCellValue("Sheet1", "F1", "Phone")
	//f.SetCellValue("Sheet1", "G1", "Detail")
	//f.SetCellValue("Sheet1", "H1", "Total")
	//
	//
	//for index, item := range lo {
	//	t := time.Unix(item.LastUpdate, 0)
	//	strDate := t.Format(time.UnixDate)
	//	row := strconv.Itoa(index + 2)
	//
	//	f.SetCellValue("Sheet1", "A"+row, item.Id)
	//	f.SetCellValue("Sheet1", "B"+row, item.Name)
	//	f.SetColWidth("Sheet1", "B", "B", 20)
	//	f.SetCellValue("Sheet1", "C"+row, item.Address)
	//	f.SetColWidth("Sheet1", "C", "C", 40)
	//	f.SetCellValue("Sheet1", "D"+row, strDate)
	//	f.SetColWidth("Sheet1", "D", "D", 40)
	//	f.SetCellValue("Sheet1", "E"+row, item.CurrentStatus)
	//	f.SetColWidth("Sheet1", "E", "E", 20)
	//	f.SetCellValue("Sheet1", "F"+row, item.Number)
	//	f.SetColWidth("Sheet1", "F", "F", 30)
	//	f.SetCellValue("Sheet1", "H"+row, item.Total)
	//	f.SetColWidth("Sheet1", "H", "H", 20)
	//
	//	detail := ""
	//	for _, i := range item.Detail {
	//		detail += "Name Product : " + i.Name + " \n" + "Is Sale : "
	//		if i.IsSale == true {
	//			detail += "Yes "
	//		} else {
	//			detail += "No "
	//		}
	//		detail += "\n"
	//		detail += "Price : " + fmt.Sprintf("%f", i.Price) + "\n"
	//		detail += "Quantity : " + strconv.Itoa(i.Quantity) + "\n"
	//	}
	//
	//	f.SetCellValue("Sheet1", "G"+row, detail)
	//	f.SetColWidth("Sheet1", "G", "G", 60)
	//	f.SetRowHeight("Sheet1", index+2, float64(50*len(item.Detail)))
	//}
	//
	//if err := f.SaveAs("ExOrder.xlsx"); err != nil {
	//	fmt.Println(err)
	//}
	//
	//return "ExOrder.xlsx", nil
	return "", nil
}

func (this *Statistics) ExImport(startTime, endTime int64) (string, error) {
	//li := []responses.ImportRes{}
	//
	//results, err := db.Query(queries.GetListImport())
	//if err != nil {
	//	return "", err
	//}
	//
	//for results.Next() {
	//	o := responses.ImportRes{}
	//	err = results.Scan(&o.Id, &o.LastUpDate, &o.Total, &o.NameAdmin, &o.VendorName)
	//	if err != nil {
	//		return "", err
	//	}
	//	li = append(li, o)
	//}
	//
	//for index, item := range li {
	//	lod, err := getImportDetail(strconv.Itoa(item.Id))
	//	if err != nil {
	//		return "", err
	//	}
	//	li[index].Detail = lod
	//}
	//
	//f := excelize.NewFile()
	//sytes, _ := f.NewStyle(`{"alignment":{"vertical":"top"}}`)
	//f.SetCellStyle("Sheet1", "E2", "E"+strconv.Itoa(len(li)+1), sytes)
	//
	//f.SetCellValue("Sheet1", "A1", "Id")
	//f.SetCellValue("Sheet1", "B1", "Vendor Name")
	//f.SetCellValue("Sheet1", "C1", "Import Time")
	//f.SetCellValue("Sheet1", "D1", "Importer")
	//f.SetCellValue("Sheet1", "E1", "Detail")
	//f.SetCellValue("Sheet1", "F1", "Total")
	//
	//for index, item := range li {
	//	t := time.Unix(item.LastUpDate, 0)
	//	strDate := t.Format(time.UnixDate)
	//	row := strconv.Itoa(index + 2)
	//
	//	f.SetCellValue("Sheet1", "A"+row, item.Id)
	//	f.SetCellValue("Sheet1", "B"+row, item.VendorName)
	//	f.SetColWidth("Sheet1", "B", "B", 40)
	//	f.SetCellValue("Sheet1", "C"+row, strDate)
	//	f.SetColWidth("Sheet1", "C", "C", 40)
	//	f.SetCellValue("Sheet1", "D"+row, item.NameAdmin)
	//	f.SetColWidth("Sheet1", "D", "D", 40)
	//	f.SetCellValue("Sheet1", "F"+row, item.Total)
	//	f.SetColWidth("Sheet1", "F", "F", 30)
	//
	//	detail := ""
	//	for _, i := range item.Detail {
	//		detail += "Name Product : " + i.Name + " \n"
	//		detail += "Price : " + fmt.Sprintf("%f", i.Price) + "\n"
	//		detail += "Quantity : " + strconv.Itoa(i.Quantity) + "\n"
	//	}
	//
	//	f.SetCellValue("Sheet1", "E"+row, detail)
	//	f.SetColWidth("Sheet1", "E", "E", 60)
	//	f.SetRowHeight("Sheet1", index+2, float64(50*len(item.Detail)))
	//}
	//
	//if err := f.SaveAs("ExImport.xlsx"); err != nil {
	//	fmt.Println(err)
	//}
	//
	//return "ExImport.xlsx", nil
	return "", nil
}

func (this *Statistics) GetTopPrice() (map[string]float64, error) {
	uPid := make(map[int]int)
	cur, err := getOrderDetailDocument().Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err.Error(), "err.Error() models/statistics.go:254")
		return map[string]float64{}, err
	}

	for cur.Next(context.TODO()) {
		var elem InsOrderDetail
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/statistics.go:260")
			return map[string]float64{}, err
		}

		if val, ok := uPid[elem.ProductId]; ok {
			uPid[elem.ProductId] = val + elem.Quantity
		} else {
			uPid[elem.ProductId] = elem.Quantity
		}
	}

	if err := cur.Err(); err != nil {
		log.Println(err.Error(), "err.Error() models/statistics.go:272")
		return map[string]float64{}, err
	}
	res := make(map[string]float64)
	for key, item := range uPid {
		p, err := (&Product{}).GetProduct(key)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/statistics.go:279")
			continue
		}
		res[p.Name] = float64(item) * p.Price
	}

	return res, nil
}

func (this *Statistics) GetTopByMonth() (map[string]float64, error) {
	cur, err := getOrderDocument().Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err.Error(), "err.Error() models/statistics.go:291")
		return map[string]float64{}, err
	}
	res := make(map[string]float64)

	for cur.Next(context.TODO()) {
		var elem Order
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/statistics.go:299")
			return map[string]float64{}, err
		}

		t := time.Unix(elem.LastUpdate, 0)
		m := t.Month()
		if val, ok := res[m.String()]; ok {
			res[m.String()] = val + elem.Total
		} else {
			res[m.String()] = elem.Total
		}
	}

	if err := cur.Err(); err != nil {
		log.Println(err.Error(), "err.Error() models/statistics.go:311")
		return map[string]float64{}, err
	}

	return res, nil
}
