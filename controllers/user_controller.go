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

type UserController struct {
	beego.Controller
}

//@Title Update User
//@Description Update User
//@Summary sửa thông tin người dùng
// @Params token header string true "Token"
// @Params data body requests.RequestUser true "data"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [put]
func (this *UserController) UpdateUser() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "user" {
		log.Println("controllers/user_controller.go:29 , typeid is not user ")
		this.Data["json"] = responses.ResponseBool{
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	uid, _ := strconv.Atoi(this.Ctx.Request.Header.Get("id"))
	req := requests.RequestUser{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println("controllers/user_controller.go:39 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	err = services.UpdateUser(req, uid)
	if err != nil {
		log.Println("controllers/user_controller.go:48v ", err)
		this.Data["json"] = responses.ResponseBool{
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error:      responses.NewErr(responses.Success),
	}
}
