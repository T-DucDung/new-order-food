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

type CartController struct {
	beego.Controller
}

//@Title Get List Item In Cart
//@Description Get List Item In Cart
//@Summary Lấy danh sách sản phẩm trong giỏ hàng
// @Params token header string true "Token"
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router / [get]
func (this *CartController) GetListItemInCart() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "user" {
		log.Println("controllers/cart_controller.go:25 , typeid is not user ")
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	uid, _ := strconv.Atoi(this.Ctx.Request.Header.Get("id"))
	log.Println("uid : ", uid)
	lc, err := services.Get(uid)
	if err != nil {
		log.Println("controllers/cart_controller.go:38 ", err)
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

//@Title Set Item Into Cart
//@Description Set Item InTo Cart
//@Summary thêm sản phẩm vào giỏ hàng
// @Params token header string true "Token"
// @Params data body requests.RequestCart true "ProductId and quantity"
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router / [post]
func (this *CartController) SetItem() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "user" {
		log.Println("controllers/cart_controller.go:66 , typeid is not user ")
		this.Data["json"] = responses.ResponseBool{
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	uid, _ := strconv.Atoi(this.Ctx.Request.Header.Get("id"))
	req := requests.RequestCart{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println("controllers/cart_controller.go:77 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	err = services.Set(uid, req.ProductId, req.Quantity)
	if err != nil {
		log.Println("controllers/cart_controller.go:86 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error:      responses.NewErr(responses.Success),
	}
}

//@Title Delete Item In Cart
//@Description Delete Item In Cart
//@Summary xóa sản phẩm trong giỏ hàng
// @Params token header string true "Token"
// @Params pid query int true "ProductId"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [delete]
func (this *CartController) RemoveItem() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "user" {
		log.Println("controllers/cart_controller.go:151 , typeid is not user ")
		this.Data["json"] = responses.ResponseBool{
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	uid, _ := strconv.Atoi(this.Ctx.Request.Header.Get("id"))
	pid, err := this.GetInt("pid", -1)
	if err != nil {
		log.Println("controllers/cart_controller.go:159 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}

	err = services.RemoveItem(uid, pid)
	if err != nil {
		log.Println("controllers/cart_controller.go:169 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error:      responses.NewErr(responses.Success),
	}
}
