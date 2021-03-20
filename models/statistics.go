package models

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"new-order-food/queries"
	"new-order-food/responses"
	"strconv"
	"time"
)

type Statistics struct {
}

func (this *Statistics) TotalRevenu() (float32, error) {
	var total float32
	err = db.QueryRow(queries.GetTotalRevenu()).Scan(&total)
	if err != nil {
		return -1, err
	}
	return total, nil
}

func (this *Statistics) TotalOrder() (int, error) {
	var total int
	err = db.QueryRow(queries.GetTotalOrder()).Scan(&total)
	if err != nil {
		return -1, err
	}
	return total, nil
}

func (this *Statistics) TopSale(num int) ([]responses.ProductRes, error) {
	p := Product{}
	return p.GetListProduct(queries.GetTopSale(num))
}

func (this *Statistics) TotalNewAcc() (int, error) {
	var total int
	err = db.QueryRow(queries.GetTotalAcc()).Scan(&total)
	if err != nil {
		return -1, err
	}
	return total, nil
}

func (this *Statistics) ExOrder(startTime, endTime int64) (string, error) {
	lo := []responses.OrderRes{}

	results, err := db.Query(queries.ExOrder(startTime, endTime))
	if err != nil {
		return "", err
	}

	for results.Next() {
		o := responses.OrderRes{}
		status := 0
		err = results.Scan(&o.Id, &o.Name, &o.Number, &o.Address, &o.Total, &status, &o.LastUpdate)
		if err != nil {
			return "", err
		}
		o.CurrentStatus = mapStatus[status]
		lo = append(lo, o)
	}

	for index, item := range lo {
		lod, err := getOrderDetail(strconv.Itoa(item.Id))
		if err != nil {
			return "", err
		}
		lo[index].Detail = lod
	}
	f := excelize.NewFile()
	sytes, _ := f.NewStyle(`{"alignment":{"vertical":"top"}}`)
	f.SetCellStyle("Sheet1", "G2", "G"+strconv.Itoa(len(lo)+1), sytes)

	f.SetCellValue("Sheet1", "A1", "Order Id")
	f.SetCellValue("Sheet1", "B1", "Cus Name")
	f.SetCellValue("Sheet1", "C1", "Address")
	f.SetCellValue("Sheet1", "D1", "Time Order")
	f.SetCellValue("Sheet1", "E1", "Order's State")
	f.SetCellValue("Sheet1", "F1", "Phone")
	f.SetCellValue("Sheet1", "G1", "Detail")

	for index, item := range lo {
		t := time.Unix(item.LastUpdate, 0)
		strDate := t.Format(time.UnixDate)
		row := strconv.Itoa(index + 2)

		f.SetCellValue("Sheet1", "A"+row, item.Id)
		f.SetCellValue("Sheet1", "B"+row, item.Name)
		f.SetColWidth("Sheet1", "B", "B", 20)
		f.SetCellValue("Sheet1", "C"+row, item.Address)
		f.SetColWidth("Sheet1", "C", "C", 40)
		f.SetCellValue("Sheet1", "D"+row, strDate)
		f.SetColWidth("Sheet1", "D", "D", 40)
		f.SetCellValue("Sheet1", "E"+row, item.CurrentStatus)
		f.SetColWidth("Sheet1", "E", "E", 20)
		f.SetCellValue("Sheet1", "F"+row, item.Number)
		f.SetColWidth("Sheet1", "F", "F", 30)

		detail := ""
		for _, i := range item.Detail {
			detail += "Name Product : " + i.Name + " \n" + "Is Sale : "
			if i.IsSale == true {
				detail += "Yes "
			} else {
				detail += "No "
			}
			detail += "\n"
			detail += "Price : " + fmt.Sprintf("%f", i.Price) + "\n"
			detail += "Quantity : " + strconv.Itoa(i.Quantity) + "\n"
		}

		f.SetCellValue("Sheet1", "G"+row, detail)
		f.SetColWidth("Sheet1", "G", "G", 60)
		f.SetRowHeight("Sheet1", index+2, float64(50*len(item.Detail)))
	}

	if err := f.SaveAs("ExOrder.xlsx"); err != nil {
		fmt.Println(err)
	}

	return "ExOrder.xlsx", nil
}

func (this *Statistics) ExImport(startTime, endTime int64) (string, error)  {
	li := []responses.ImportRes{}

	results, err := db.Query(queries.GetListImport())
	if err != nil {
		return "", err
	}

	for results.Next() {
		o := responses.ImportRes{}
		err = results.Scan(&o.Id, &o.LastUpDate, &o.Total, &o.NameAdmin, &o.VendorName)
		if err != nil {
			return "", err
		}
		li = append(li, o)
	}

	for index, item := range li {
		lod, err := getImportDetail(strconv.Itoa(item.Id))
		if err != nil {
			return "", err
		}
		li[index].Detail = lod
	}

	f := excelize.NewFile()
	sytes, _ := f.NewStyle(`{"alignment":{"vertical":"top"}}`)
	f.SetCellStyle("Sheet1", "E2", "E"+strconv.Itoa(len(li)+1), sytes)

	f.SetCellValue("Sheet1", "A1", "Id")
	f.SetCellValue("Sheet1", "B1", "Vendor Name")
	f.SetCellValue("Sheet1", "C1", "Import Time")
	f.SetCellValue("Sheet1", "D1", "Importer")
	f.SetCellValue("Sheet1", "E1", "Detail")
	f.SetCellValue("Sheet1", "F1", "Total")

	for index, item := range li {
		t := time.Unix(item.LastUpDate, 0)
		strDate := t.Format(time.UnixDate)
		row := strconv.Itoa(index + 2)

		f.SetCellValue("Sheet1", "A"+row, item.Id)
		f.SetCellValue("Sheet1", "B"+row, item.VendorName)
		f.SetColWidth("Sheet1", "B", "B", 40)
		f.SetCellValue("Sheet1", "C"+row, strDate)
		f.SetColWidth("Sheet1", "C", "C", 40)
		f.SetCellValue("Sheet1", "D"+row, item.NameAdmin)
		f.SetColWidth("Sheet1", "D", "D", 40)
		f.SetCellValue("Sheet1", "F"+row, item.Total)
		f.SetColWidth("Sheet1", "F", "F", 30)

		detail := ""
		for _, i := range item.Detail {
			detail += "Name Product : " + i.Name + " \n"
			detail += "Price : " + fmt.Sprintf("%f", i.Price) + "\n"
			detail += "Quantity : " + strconv.Itoa(i.Quantity) + "\n"
		}

		f.SetCellValue("Sheet1", "E"+row, detail)
		f.SetColWidth("Sheet1", "E", "E", 60)
		f.SetRowHeight("Sheet1", index+2, float64(50*len(item.Detail)))
	}

	if err := f.SaveAs("ExImport.xlsx"); err != nil {
		fmt.Println(err)
	}

	return "ExImport.xlsx", nil
}
