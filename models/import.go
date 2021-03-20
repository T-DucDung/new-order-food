package models

import (
	"new-order-food/queries"
	"new-order-food/responses"
	"strconv"
	"time"
)

type Import struct {
	Id         int            `json:"id" xml:"id"`
	VendorId   int            `json:"vendor_id" xml:"vendor_id"`
	LastUpDate int64          `json:"last_up_date" xml:"last_up_date"`
	IdAdmin    int            `json:"id_admin" xml:"id_admin"`
	Total      float32        `json:"total" xml:"total"`
	Detail     []ImportDetail `json:"detail" xml:"detail"`
}

type ImportDetail struct {
	ProductId int     `json:"product_id" xml:"product_id"`
	Unit      string  `json:"unit" xml:"unit"`
	Quantity  int     `json:"quantity" xml:"quantity"`
	Price     float32 `json:"price" xml:"price"`
}

func (this *Import) Import() error {
	data, err := db.Prepare("insert into Import (IdAdmin,LastUpDate,Total,VendorId) VALUES(?, ?, ?, ?);")
	if err != nil {
		return err
	}
	val, err := data.Exec(this.IdAdmin, time.Now().Unix(), this.Total, this.VendorId)
	if err != nil {
		return err
	}

	id, err := val.LastInsertId()
	p := Product{}

	for _, item := range this.Detail {
		data, err = db.Prepare("insert into ImportDetail (IdImport,ProductId,Quantity,Price) VALUES(?, ?, ?, ?);")
		if err != nil {
			return err
		}
		_, err = data.Exec(id, item.ProductId, item.Quantity, item.Price)
		if err != nil {
			return err
		}
		c, _ := p.CheckExist(item.ProductId)
		if c == true {
			err = p.UpdateRemaining(strconv.Itoa(item.ProductId), item.Quantity, 0)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *Import) GetAllImport() ([]responses.ImportRes, error) {
	li := []responses.ImportRes{}

	results, err := db.Query(queries.GetListImport())
	if err != nil {
		return nil, err
	}

	for results.Next() {
		o := responses.ImportRes{}
		err = results.Scan(&o.Id, &o.LastUpDate, &o.Total, &o.NameAdmin, &o.VendorName)
		if err != nil {
			return nil, err
		}
		li = append(li, o)
	}

	for index, item := range li {
		lod, err := getImportDetail(strconv.Itoa(item.Id))
		if err != nil {
			return nil, err
		}
		li[index].Detail = lod
	}

	return li, nil
}

func getImportDetail(id string) ([]responses.ImportDetailRes , error){
	lid := []responses.ImportDetailRes{}

	results, err := db.Query(queries.GetListImportDetail(id))
	if err != nil {
		return nil, err
	}

	for results.Next() {
		o := responses.ImportDetailRes{}
		err = results.Scan(&o.Name, &o.Quantity, &o.Price)
		if err != nil {
			return nil, err
		}
		lid = append(lid, o)
	}

	return lid, nil
}
