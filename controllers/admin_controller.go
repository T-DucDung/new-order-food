package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"new-order-food/requests"
	"new-order-food/responses"
	"new-order-food/services"
)

type AdminController struct {
	beego.Controller
}

//@Title Create Account
//@Description Create Account
//@Summary Tạo mới tài khoản
// @Params token header string true "Token"
// @Param req body requests.RequestCreateAccount true "req"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [post]
func (this *AdminController) CreateAccount() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/admin_controller.go:28 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	req := requests.RequestCreateAccount{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println("controllers/admin_controller.go:37 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.CreateAccount(req)
	if err != nil {
		log.Println("controllers/admin_controller.go:45 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}
