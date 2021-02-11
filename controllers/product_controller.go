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
// @Params token header string true "Token"
// @Param data body requests.RequestCreateProduct true "product"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [post]
func (this *ProductController) CreateProduct() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/product_controller.go:29 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	product := requests.RequestCreateProduct{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &product)
	if err != nil {
		log.Println("controllers/product_controller.go:38 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.CreateProduct(product)
	if err != nil {
		log.Println("controllers/product_controller.go:46 ", err)
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
		log.Println("controllers/product_controller.go:70 ", err)
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
// @Param cateid query string false "cate id"
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router /list [get]
func (this *ProductController) GetListProduct() {
	defer this.ServeJSON()
	cateid := this.GetString("cateid")
	log.Println("cate id : ", cateid)
	lp, err := services.GetListProduct(cateid)
	if err != nil {
		log.Println("controllers/product_controller.go:96 ", err)
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

//@Title Update Product
//@Description Update Product
//@Summary sửa một sản phẩm
// @Params token header string true "Token"
// @Param data body models.Product true "product"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router / [put]
func (this *ProductController) UpDateProduct()  {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/product_controller.go:123 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	product := models.Product{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &product)
	if err != nil {
		log.Println("controllers/product_controller.go:132 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.UpDateProduct(product)
	if err != nil {
		log.Println("controllers/product_controller.go:140 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}
