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

//@Title Register
//@Description Register
//@Summary Tạo mới tài khoản
// @Param req body requests.RequestRegister true "req"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router /register [post]
func (this *AccountController) Register() {
	defer this.ServeJSON()
	req := requests.RequestRegister{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println("controllers/account_controller.go:62 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.Register(req)
	if err != nil {
		log.Println("controllers/account_controller.go:70 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Register For Admin
//@Description Register For Admin
//@Summary Tạo mới tài khoản
// @Params token header string true "Token"
// @Param req body requests.RequestRegisterForAdmin true "req"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router /auth [post]
func (this *AccountController) RegisterForAdmin() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/account_controller.go:93 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	req := requests.RequestRegisterForAdmin{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println("controllers/account_controller.go:102 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.RegisterForAdmin(req)
	if err != nil {
		log.Println("controllers/account_controller.go:110 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}
