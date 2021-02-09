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
// @Param data body models.Category true "category"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router / [put]
func (this *CategoryController) UpDateCategory()  {
	defer this.ServeJSON()
	cate := models.Category{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &cate)
	if err != nil {
		log.Println("controllers/category_controller.go:53 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.UpdateCategory(cate)
	if err != nil {
		log.Println("controllers/category_controller.go:61 ", err)
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
// @Param name body string true "name categorry"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [post]
func (this *CategoryController) CreateCategory() {
	defer this.ServeJSON()
	var nameCate string
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &nameCate)
	log.Println("nameCate ", nameCate)
	if err != nil {
		log.Println("controllers/category_controller.go:85 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.CreateCategory(nameCate)
	if err != nil {
		log.Println("controllers/category_controller.go:93 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}
