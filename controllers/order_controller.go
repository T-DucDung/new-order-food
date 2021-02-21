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
	if idtype != "admin" {
		log.Println("controllers/order_controller.go:28 , typeid is not user ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	id := this.Ctx.Request.Header.Get("id")
	uid,_ := strconv.Atoi(id)

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
