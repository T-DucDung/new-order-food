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
			Error: responses.NewErr(responses.UnSuccess),
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

//@Title Get List User
//@Description Get List User
//@Summary Lấy một danh sách người dùng
// @Params token header string true "Token"
// @Param page query int false "page"
// @Param size query int false "size"
// @Param status query string false "status"
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router / [get]
func (this *AdminController) GetAllUser() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/admin_controller.go:69 , typeid is not admin ")
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}

	status := this.GetString("status")
	page, err := this.GetInt("page", 1)
	size, err := this.GetInt("size", 10)

	lu, count, err := services.GetAllUser((page-1)*size, size, status)
	if err != nil {
		log.Println("controllers/admin_controller.go:82 ", err)
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseArray{
		Data:       lu,
		TotalCount: count,
		Error:      responses.NewErr(responses.Success),
	}
}

//@Title Update Status
//@Description Update Status
//@Summary cập nhật lại trạng thái người dùng
// @Params token header string true "Token"
// @Param data body requests.RequestUpdateStatus true "data"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [put]
func (this *AdminController) UpdateStatus() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/admin_controller.go:110 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	req := requests.RequestUpdateStatus{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)

	err := services.UpdateStatus(req)
	if err != nil {
		log.Println("controllers/admin_controller.go:123 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
	return
}
