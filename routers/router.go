package routers

import (
	"github.com/astaxie/beego"
	"new-order-food/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/product",
			beego.NSInclude(
				&controllers.ProductController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
