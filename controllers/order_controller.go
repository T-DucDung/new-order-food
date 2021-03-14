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

type OrderController struct {
	beego.Controller
}

//@Title Pay Order
//@Description Pay Order
//@Summary thanh toán đơn hàng
// @Params token header string true "Token"
// @Param data body requests.RequestOrder true "order"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [post]
func (this *OrderController) PayOrder() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "user" {
		log.Println("controllers/order_controller.go:28 , typeid is not user ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	id := this.Ctx.Request.Header.Get("id")
	uid, _ := strconv.Atoi(id)

	order := requests.RequestOrder{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &order)
	if err != nil {
		log.Println("controllers/order_controller.go:41 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.PayOrder(order, uid)
	if err != nil {
		log.Println("controllers/order_controller.go:49 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Get List Order
//@Description Get List Order
//@Summary Lấy danh sách đặt hàng
// @Params token header string true "Token"
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router /list [get]
func (this *OrderController) GetListOrder() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "user" {
		log.Println("controllers/order_controller.go:71 , typeid is not user ")
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	uid := this.Ctx.Request.Header.Get("id")

	lo, err := services.GetListOrder(uid)
	if err != nil {
		log.Println("controllers/order_controller.go:83", err)
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseArray{
		Data:       lo,
		TotalCount: len(lo),
		Error:      responses.NewErr(responses.Success),
	}
}

//@Title Get List Order
//@Description Get List Order
//@Summary Lấy danh sách đặt hàng
// @Params token header string true "Token"
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router /listAll [get]
func (this *OrderController) GetListOrderForAdmin() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/order_controller.go:110 , typeid is not admin ")
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}

	lo, err := services.GetListOrderForAdmin()
	if err != nil {
		log.Println("controllers/order_controller.go:121 ", err)
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseArray{
		Data:       lo,
		TotalCount: len(lo),
		Error:      responses.NewErr(responses.Success),
	}
}

//@Title Update Category
//@Description Update Category
//@Summary sửa đơn hàng
// @Params token header string true "Token"
// @Param id path string true "id"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router /:id [put]
func (this *OrderController) UpDateOrder() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/order_controller.go:148 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	id := this.GetString(":id")
	log.Println("id ", id)
	err := services.UpDateOrder(id)
	if err != nil {
		log.Println("controllers/order_controller.go:166 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}
