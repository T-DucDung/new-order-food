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

type VendorController struct {
	beego.Controller
}

//@Title Get List Vendor
//@Description Get List Vendor
//@Summary Lấy danh sách nhà cung cấp
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router / [get]
func (this *VendorController) GetListVendor() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/vendor_controller.go:26 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	lv, err := services.GetListVendor()
	if err != nil {
		log.Println("controllers/category_controller.go:35 ", err)
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseArray{
		Data:       lv,
		TotalCount: len(lv),
		Error:      responses.NewErr(responses.Success),
	}
}

//@Title Update Vendor
//@Description Update Vendor
//@Summary sửa một nhà cung cấp
// @Params token header string true "Token"
// @Param data body models.Vendor true "category"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [put]
func (this *VendorController) UpDateVendor() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/vendor_controller.go:62 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	vendor := models.Vendor{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &vendor)
	if err != nil {
		log.Println("controllers/vendor_controller.go:72 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.UpdateVendor(vendor)
	if err != nil {
		log.Println("controllers/vendor_controller.go:79 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Create Vendor
//@Description Create Vendor
//@Summary Tạo một nhà cung cấp
// @Params token header string true "Token"
// @Param data body requests.RequestVendor true "data"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [post]
func (this *VendorController) CreateCategory() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/vendor_controller.go:102 , typeid is not admin ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	vendor := requests.RequestVendor{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &vendor)
	log.Println("vendor ", vendor)
	if err != nil {
		log.Println("controllers/vendor_controller.go:113 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.CreateVendor(vendor)
	if err != nil {
		log.Println("controllers/vendor_controller.go:121 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}
