package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"new-order-food/models"
	"new-order-food/requests"
	"new-order-food/responses"
	"new-order-food/services"
)

type UserController struct {
	beego.Controller
}


//@Title Get List User
//@Description Get List User
//@Summary Lấy một danh sách sản phẩm
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router /list [get]
func (this *UserController) GetListUser() {
	defer this.ServeJSON()

	lp, err := services.GetListUser()
	if err != nil {
		log.Println("controllers/user_controller.go:29 ", err)
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseArray{
		Data:       lp,
		TotalCount: len(lp),
		Error:      responses.NewErr(responses.Success),
	}
}