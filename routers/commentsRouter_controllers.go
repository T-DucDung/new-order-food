package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new-order-food/controllers:AccountController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:CategoryController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetListCategory",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:CategoryController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "UpDateCategory",
            Router: "/admin",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:CategoryController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "CreateCategory",
            Router: "/admin",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:ProductController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:ProductController"],
        beego.ControllerComments{
            Method: "CreateProduct",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:ProductController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetProduct",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:ProductController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:ProductController"],
        beego.ControllerComments{
            Method: "UpDateProduct",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:ProductController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetListProduct",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
