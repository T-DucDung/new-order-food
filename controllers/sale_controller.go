package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"new-order-food/requests"
	"new-order-food/responses"
	"new-order-food/services"
	"strconv"
)

type SaleController struct {
	beego.Controller
}

//@Title Get List Sale
//@Description Get List Sale
// @Params token header string true "Token"
//@Summary Lấy danh sách sale
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router / [get]
func (this *SaleController) GetListSale() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/sale_controller.go:25 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	lc, err := services.GetSale()
	if err != nil {
		log.Println("controllers/sale_controller.go:34 ", err)
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseArray{
		Data:       lc,
		TotalCount: len(lc),
		Error:      responses.NewErr(responses.Success),
	}
}

//@Title Create Sale
//@Description Create sale
//@Summary Tạo sale
// @Params token header string true "Token"
// @Param data body requests.RequestSale true "data"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [post]
func (this *SaleController) CreateSale() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/sale_controller.go:64 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	uid, _ := strconv.Atoi(this.Ctx.Request.Header.Get("id"))

	req := requests.RequestSale{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)

	if err != nil {
		log.Println("controllers/sale_controller.go:76 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.CreateSale(req, uid)
	if err != nil {
		log.Println("controllers/sale_controller.go:84 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Update Sale
//@Description Update Sale
//@Summary sửa sale
// @Params token header string true "Token"
// @Param id path string true "id"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router /:id [put]
func (this *SaleController) UpDateSale() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/sale_controller.go:107 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	id, err := this.GetInt(":id", -1)
	if err != nil {
		log.Println("controllers/sale_controller.go:115 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	log.Println("id ", id)
	err = services.UpdateSale(id)
	if err != nil {
		log.Println("controllers/sale_controller.go:124 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}
