package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"new-order-food/responses"
	"new-order-food/services"
	"time"
)

type StatisticController struct {
	beego.Controller
}

//@Title Get Total Revenu
//@Description Get Total Revenu
//@Summary Lấy tổng thu nhập
// @Params token header string true "Token"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router /total-revenu [get]
func (this *StatisticController) GetTotalRevenu() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/statistic_controller.go:26 , typeid is not admin ")
		this.Data["json"] = responses.ResponseSingle{
			Data:  nil,
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	total, err := services.TotalRevenu()
	if err != nil {
		log.Println("controllers/statistic_controller.go:35 ", err)
		this.Data["json"] = responses.ResponseSingle{
			Data:  nil,
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseSingle{
		Data:  total,
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Get Total Order
//@Description Get Total Order
//@Summary Lấy tổng order
// @Params token header string true "Token"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router /total-order [get]
func (this *StatisticController) GetTotalOrder() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/statistic_controller.go:58 , typeid is not admin ")
		this.Data["json"] = responses.ResponseSingle{
			Data:  nil,
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	total, err := services.TotalOrder()
	if err != nil {
		log.Println("controllers/statistic_controller.go:67 ", err)
		this.Data["json"] = responses.ResponseSingle{
			Data:  nil,
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseSingle{
		Data:  total,
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Get Top Sale
//@Description Get Top Sale
//@Summary Lấy sản phẩm bán chạy
// @Params num query int false "num"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router /total-top-sale [get]
func (this *StatisticController) GetTopSale() {
	defer this.ServeJSON()
	num, _ := this.GetInt("num", 5)

	total, err := services.TopSale(num)
	if err != nil {
		log.Println("controllers/statistic_controller.go:104 ", err)
		this.Data["json"] = responses.ResponseSingle{
			Data:  nil,
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseSingle{
		Data:  total,
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Get Excel
//@Description Get Excel
//@Summary Xuất excel
// @Params token header string true "Token"
// @Params startTime query int64 false "Start Time"
// @Params endTime query int64 false "End Time"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router /ex-order [get]
func (this *StatisticController) ExOrder() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/statistic_controller.go:119 , typeid is not admin ")
		this.Data["json"] = responses.ResponseSingle{
			Data:  nil,
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	startTime, _ := this.GetInt64("startTime", 1612026000)
	endTime, _ := this.GetInt64("endTime", time.Now().Unix())

	f, err := services.ExOrder(startTime, endTime)
	if err != nil {
		log.Println("controllers/statistic_controller.go:133 ", err)
		this.Data["json"] = responses.ResponseSingle{
			Data:  nil,
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	ts := time.Unix(startTime, 0)
	te := time.Unix(endTime, 0)

	this.Ctx.Output.SetStatus(200)
	this.Ctx.Output.Download(f, "ExOrder-"+fmt.Sprintf("%02d-%02d-%d", ts.Day(), ts.Month(), ts.Year())+"-"+fmt.Sprintf("%02d-%02d-%d", te.Day(), te.Month(), te.Year())+".xlsx")
}

//@Title Get Total Account
//@Description Get Total Account
//@Summary Lấy tổng tài khoản khách hàng
// @Params token header string true "Token"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router /total-acc [get]
func (this *StatisticController) GetTotalAccount() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/statistic_controller.go:128 , typeid is not admin ")
		this.Data["json"] = responses.ResponseSingle{
			Data:  nil,
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	total, err := services.TotalNewAcc()
	if err != nil {
		log.Println("controllers/statistic_controller.go:138 ", err)
		this.Data["json"] = responses.ResponseSingle{
			Data:  nil,
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseSingle{
		Data:  total,
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Get Excel Import
//@Description Get Excel Import
//@Summary Xuất excel nhập
// @Params token header string true "Token"
// @Params startTime query int64 false "Start Time"
// @Params endTime query int64 false "End Time"
//@Success 200 {object} responses.ResponseSingle
//@Failure 404 {object} responses.ResponseSingle
//@router /ex-import [get]
func (this *StatisticController) ExImport() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "admin" {
		log.Println("controllers/statistic_controller.go:195, typeid is not admin ")
		this.Data["json"] = responses.ResponseSingle{
			Data:  nil,
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	startTime, _ := this.GetInt64("startTime", 1612026000)
	endTime, _ := this.GetInt64("endTime", time.Now().Unix())

	f, err := services.ExImport(startTime, endTime)
	if err != nil {
		log.Println("controllers/statistic_controller.go:208 ", err)
		this.Data["json"] = responses.ResponseSingle{
			Data:  nil,
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	ts := time.Unix(startTime, 0)
	te := time.Unix(endTime, 0)

	this.Ctx.Output.SetStatus(200)
	this.Ctx.Output.Download(f, "ExImport-"+fmt.Sprintf("%02d-%02d-%d", ts.Day(), ts.Month(), ts.Year())+"-"+fmt.Sprintf("%02d-%02d-%d", te.Day(), te.Month(), te.Year())+".xlsx")
}
