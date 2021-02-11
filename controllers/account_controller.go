package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"new-order-food/requests"
	"new-order-food/responses"
	"new-order-food/services"
)

type AccountController struct {
	beego.Controller
}

//@Title Login
//@Description Login
//@Summary đăng nhập
// @Param data body requests.RequestLogin true "login"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router / [post]
func (this *AccountController) Login() {
	defer this.ServeJSON()
	data := requests.RequestLogin{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	if err != nil {
		log.Println("controllers/account_controller.go:28 ", err)
		this.Data["json"] = responses.ResponseSingle{
			Data:  "",
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	token, err := services.Login(data)
	if err != nil {
		log.Println("controllers/account_controller.go:37 ", err)
		this.Data["json"] = responses.ResponseSingle{
			Data:  "",
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseSingle{
		Data:  token,
		Error: responses.NewErr(responses.Success),
	}
}
