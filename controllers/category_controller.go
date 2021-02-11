package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"new-order-food/models"
	"new-order-food/responses"
	"new-order-food/services"
)

type CategoryController struct {
	beego.Controller
}

//@Title Get List Category
//@Description Get List Category
//@Summary Lấy danh sách loại sản phẩm
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router / [get]
func (this *CategoryController) GetListCategory() {
	defer this.ServeJSON()
	lc, err := services.GetListCategory()
	if err != nil {
		log.Println("controllers/category_controller.go:24 ", err)
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseArray{
		Data:  lc,
		TotalCount: len(lc),
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Update Category
//@Description Update Category
//@Summary sửa một loại sản phẩm
// @Params token header string true "Token"
// @Param data body models.Category true "category"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router /admin [put]
func (this *CategoryController) UpDateCategory()  {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/category_controller.go:53, typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	cate := models.Category{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &cate)
	if err != nil {
		log.Println("controllers/category_controller.go:62 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.UpdateCategory(cate)
	if err != nil {
		log.Println("controllers/category_controller.go:70 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Create Category
//@Description Create Category
//@Summary Tạo một loại sản phẩm
// @Params token header string true "Token"
// @Param name body string true "name categorry"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router /admin [post]
func (this *CategoryController) CreateCategory() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/category_controller.go:93, typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	var nameCate string
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &nameCate)
	log.Println("nameCate ", nameCate)
	if err != nil {
		log.Println("controllers/category_controller.go:103 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.CreateCategory(nameCate)
	if err != nil {
		log.Println("controllers/category_controller.go:111 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}
