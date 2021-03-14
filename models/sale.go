package models

import (
	"new-order-food/queries"
	"new-order-food/responses"
	"strconv"
	"time"
)

type SaleCampaign struct {
	Id         int                  `json:"id" xml:"id"`
	IdAdmin    int                  `json:"id_admin" xml:"id_admin"`
	LastUpDate int64                `json:"last_up_date" xml:"last_up_date"`
	Status     bool                 `json:"status" xml:"status"`
	Detail     []DetailSaleCampaign `json:"detail" xml:"detail"`
}

type DetailSaleCampaign struct {
	ProductId int     `json:"product_id" xml:"product_id"`
	SalePrice float32 `json:"sale_price" xml:"sale_price"`
}

func (this *SaleCampaign) Create() error {
	data, err := db.Prepare("insert into SaleCampaign (IdAdmin,LastUpDate,Status) VALUES(?, ?, ?);")
	if err != nil {
		return err
	}
	val, err := data.Exec(this.IdAdmin, time.Now().Unix(), this.Status)
	if err != nil {
		return err
	}

	id, err := val.LastInsertId()
	p := Product{}

	for _, item := range this.Detail {
		data, err = db.Prepare("insert into SaleCampaignDetail (CampaignId,ProductId,SalePrice) VALUES(?, ?, ?);")
		if err != nil {
			return err
		}
		_, err = data.Exec(id, item.ProductId, item.SalePrice)
		if err != nil {
			return err
		}
		c, _ := p.CheckExist(item.ProductId)
		if c == true {
			err = p.UpdateSalePrice(strconv.Itoa(item.ProductId), item.SalePrice, this.Status)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *SaleCampaign) GetList() ([]responses.SaleCampaignRes, error) {
	ls := []responses.SaleCampaignRes{}

	results, err := db.Query(queries.GetAllSale())
	if err != nil {
		return nil, err
	}

	for results.Next() {
		o := responses.SaleCampaignRes{}
		err = results.Scan(&o.Id, &o.NameAdmin, &o.LastUpDate, &o.Status)
		if err != nil {
			return nil, err
		}
		ls = append(ls, o)
	}

	for index, item := range ls {
		lod, err := getSaleDetail(strconv.Itoa(item.Id))
		if err != nil {
			return nil, err
		}
		ls[index].Detail = lod
	}

	return ls, nil
}

func getSaleDetail(id string) ([]responses.DetailSaleCampaignRes, error) {
	lsc := []responses.DetailSaleCampaignRes{}

	results, err := db.Query(queries.GetDetailSale(id))
	if err != nil {
		return nil, err
	}

	for results.Next() {
		o := responses.DetailSaleCampaignRes{}
		err = results.Scan(&o.Name, &o.SalePrice)
		if err != nil {
			return nil, err
		}
		lsc = append(lsc, o)
	}

	return lsc, nil
}

func (this *SaleCampaign) UpDateStatus() error {
	data, err := db.Prepare("UPDATE SaleCampaign as s SET s.Status = ? WHERE s.Id = ?;")
	if err != nil {
		return err
	}
	_, err = data.Exec(this.Status, this.Id)
	if err != nil {
		return err
	}

	results, err := db.Query(queries.GetIdProduct(strconv.Itoa(this.Id)))
	if err != nil {
		return err
	}
	p := Product{}
	for results.Next() {
		var id int
		err = results.Scan(&id)
		if err != nil {
			return err
		}
		c, _ := p.CheckExist(id)
		if c == true {
			err = p.UpdateSalePrice(strconv.Itoa(id), 0, this.Status)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
