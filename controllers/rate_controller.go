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

type RateController struct {
	beego.Controller
}

//@Title Set Rate
//@Description Set Rate
//@Summary đánh giá
// @Params token header string true "Token"
// @Params data body requests.RequestRate true "data"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [post]
func (this *RateController) SetRate() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "user" {
		log.Println(" , typeid is not user ")
		this.Data["json"] = responses.ResponseBool{
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	uid, _ := strconv.Atoi(this.Ctx.Request.Header.Get("id"))
	req := requests.RequestRate{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println("controllers/rate_controller.go:39 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	err = services.SetRate(uid, req.ProductId, req.Rate)
	if err != nil {
		log.Println("controllers/rate_controller.go:48 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error:      responses.NewErr(responses.Success),
	}
}