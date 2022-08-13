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
// @Param pid path int true "product id"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router /:pid [get]
func (this *ProductController) GetProduct() {
	defer this.ServeJSON()
	pid, _ := this.GetInt(":pid")
	//log.Println("product id : ", pid)
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
	return
}

//@Title Get List Product
//@Description Get List Product
//@Summary Lấy một danh sách sản phẩm
// @Param page query int false "page"
// @Param size query int false "size"
// @Param cateid query int false "cate id"
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router /list [get]
func (this *ProductController) GetListProduct() {
	defer this.ServeJSON()
	cateid, _ := this.GetInt("cateid", -1)
	page, _ := this.GetInt("page", 1)
	size, _ := this.GetInt("size", 10)
	//log.Println(page, size, cateid)
	lp, count, err := services.GetListProduct((page-1)*size, size, cateid)
	if err != nil {
		log.Println("controllers/product_controller.go:100 ", err)
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseArray{
		Data:       lp,
		TotalCount: count,
		Error:      responses.NewErr(responses.Success),
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
func (this *ProductController) UpDateProduct() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/product_controller.go:154 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	product := models.Product{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &product)
	if err != nil {
		log.Println("controllers/product_controller.go:163 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.UpdateProduct(product)
	if err != nil {
		log.Println("controllers/product_controller.go:171 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Get List Product Search
//@Description Get List Product Search
//@Summary Tìm kiếm danh sách sản phẩm
// @Param word query string true "word"
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router /search [get]
func (this *ProductController) GetListSearch() {
	defer this.ServeJSON()
	//word := this.GetString("word")
	//lp, err := services.Search(word)
	//if err != nil {
	//	log.Println("controllers/product_controller.go:167 ", err)
	//	this.Data["json"] = responses.ResponseArray{
	//		Data:       nil,
	//		TotalCount: 0,
	//		Error:      responses.NewErr(responses.UnSuccess),
	//	}
	//	return
	//}
	//this.Data["json"] = responses.ResponseArray{
	//	Data:       lp,
	//	TotalCount: len(lp),
	//	Error:      responses.NewErr(responses.Success),
	//}
}
