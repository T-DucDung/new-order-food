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

type ProductController struct {
	beego.Controller
}

//@Title Create Product
//@Description Create Product
//@Summary Tạo mới sản phẩm
// @Param data body requests.RequestCreateProduct true "product"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [post]
func (this *ProductController) CreateProduct() {
	defer this.ServeJSON()
	product := requests.RequestCreateProduct{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &product)
	if err != nil {
		log.Println("controllers/product_controller.go:26 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.CreateProduct(product)
	if err != nil {
		log.Println("controllers/product_controller.go:36 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Get Product
//@Description Get Product
//@Summary Lấy một sản phẩm
// @Param pid query string true "product id"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router / [get]
func (this *ProductController) GetProduct() {
	defer this.ServeJSON()
	pid := this.GetString("pid")
	log.Println("product id : ", pid)
	p, err := services.GetProduct(pid)
	if err != nil {
		log.Println("controllers/product_controller.go:60 ", err)
		this.Data["json"] = responses.ResponseSingle{
			Data:  nil,
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseSingle{
		Data:  p,
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Get List Product
//@Description Get List Product
//@Summary Lấy một danh sách sản phẩm
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router /list [get]
func (this *ProductController) GetListProduct() {
	defer this.ServeJSON()
	lp, err := services.GetListProduct()
	if err != nil {
		log.Println("controllers/product_controller.go:83 ", err)
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseArray{
		Data:  lp,
		TotalCount: len(lp),
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Up Date Product
//@Description Up Date Product
//@Summary sửa một sản phẩm
// @Param data body models.Product true "product"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router / [put]
func (this *ProductController) UpDateProduct()  {
	defer this.ServeJSON()
	product := models.Product{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &product)
	if err != nil {
		log.Println("controllers/product_controller.go:112 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.UpDateProduct(product)
	if err != nil {
		log.Println("controllers/product_controller.go:119 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}
