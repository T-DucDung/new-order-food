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
		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),
		beego.NSNamespace("/account",
			beego.NSInclude(
				&controllers.AccountController{},
			),
		),
		beego.NSNamespace("/cart",
			beego.NSInclude(
				&controllers.CartController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
