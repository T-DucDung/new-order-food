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

type DiscountController struct {
	beego.Controller
}

//@Title Get List Discount
//@Description Get List Discount
//@Summary Lấy danh sách giảm giá
// @Params token header string true "Token"
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router / [get]
func (this *DiscountController) GetListDiscount() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/discount_controller.go:27 , typeid is not admin ")
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	ld, err := services.GetListDiscount()
	if err != nil {
		log.Println("controllers/discount_controller.go:37 ", err)
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseArray{
		Data:       ld,
		TotalCount: len(ld),
		Error:      responses.NewErr(responses.Success),
	}
}

//@Title Update Discount
//@Description Update Discount
//@Summary sửa một giảm giá
// @Params token header string true "Token"
// @Param data body requests.RequestDis true "category"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [put]
func (this *DiscountController) UpdateDiscount() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/discount_controller.go:64 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	uid, _ := strconv.Atoi(this.Ctx.Request.Header.Get("id"))

	dis := requests.RequestDis{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &dis)
	if err != nil {
		log.Println("controllers/discount_controller.go:74 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.UpdateDiscount(dis, uid)
	if err != nil {
		log.Println("controllers/discount_controller.go:85 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Create Discount
//@Description Create Discount
//@Summary Tạo một loại giảm giá
// @Params token header string true "Token"
// @Param data body requests.RequestDis true "data"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [post]
func (this *DiscountController) CreateDiscount() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/discount_controller.go:108 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	uid, _ := strconv.Atoi(this.Ctx.Request.Header.Get("id"))

	dis := requests.RequestDis{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &dis)
	if err != nil {
		log.Println("controllers/category_controller.go:103 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.CreateDiscount(dis, uid)
	if err != nil {
		log.Println("controllers/discount_controller.go:126 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}
