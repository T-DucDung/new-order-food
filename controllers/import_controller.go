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

type ImportController struct {
	beego.Controller
}

//@Title Get List Import
//@Description Get List Import
// @Params token header string true "Token"
//@Summary Lấy danh sách nhập
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router / [get]
func (this *ImportController) GetListImport() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/import_controller.go:25 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	lc, err := services.GetAllImport()
	if err != nil {
		log.Println("controllers/import_controller.go:34 ", err)
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


//@Title Create Import
//@Description Create Import
//@Summary Tạo đơn nhập
// @Params token header string true "Token"
// @Param data body requests.RequestImport true "data"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [post]
func (this *ImportController) CreateImport() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/category_controller.go:93, typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	uid, _ := strconv.Atoi(this.Ctx.Request.Header.Get("id"))

	req := requests.RequestImport{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)

	if err != nil {
		log.Println("controllers/import_controller.go:77 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.Import(req, uid)
	if err != nil {
		log.Println("controllers/import_controller.go:85 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}
