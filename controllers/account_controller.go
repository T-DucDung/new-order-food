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
//@router /login [post]
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
	token, typeid, err := services.Login(data)
	if err != nil {
		log.Println("controllers/account_controller.go:37 ", err)
		this.Data["json"] = responses.ResponseSingle{
			Data:  "",
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseSingle{
		Data: map[string]string{
			"token":  token,
			"typeid": typeid,
		},
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

//@Title Check Exist Token)
//@Description Check Exist Token
//@Summary check tồn tại token
// @Param token query string true "token"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router / [get]
func (this *AccountController) CheckExistToken() {
	defer this.ServeJSON()
	token := this.GetString("token")
	//log.Println("token : ", token)
	r, err := services.CheckExistToken(token)
	if err != nil {
		log.Println("controllers/account_controller.go:97 ", err)
		this.Data["json"] = responses.ResponseSingle{
			Data:  nil,
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseSingle{
		Data:  r,
		Error: responses.NewErr(responses.Success),
	}
}
