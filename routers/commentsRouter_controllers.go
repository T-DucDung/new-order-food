package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new-order-food/controllers:AccountController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:AccountController"],
        beego.ControllerComments{
            Method: "CheckExistToken",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:AccountController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:AccountController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:AdminController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:AdminController"],
        beego.ControllerComments{
            Method: "CreateAccount",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:AdminController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetAllUser",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:AdminController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:AdminController"],
        beego.ControllerComments{
            Method: "UpdateStatus",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:CartController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:CartController"],
        beego.ControllerComments{
            Method: "GetListItemInCart",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:CartController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:CartController"],
        beego.ControllerComments{
            Method: "SetItem",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:CartController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:CartController"],
        beego.ControllerComments{
            Method: "RemoveItem",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
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
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:CategoryController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "CreateCategory",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:CommentController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:CommentController"],
        beego.ControllerComments{
            Method: "GetListComment",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:CommentController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:CommentController"],
        beego.ControllerComments{
            Method: "CreateComment",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:DiscountController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:DiscountController"],
        beego.ControllerComments{
            Method: "GetListDiscount",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:DiscountController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:DiscountController"],
        beego.ControllerComments{
            Method: "UpdateDiscount",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:DiscountController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:DiscountController"],
        beego.ControllerComments{
            Method: "CreateDiscount",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:OrderController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:OrderController"],
        beego.ControllerComments{
            Method: "PayOrder",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:OrderController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:OrderController"],
        beego.ControllerComments{
            Method: "UpDateOrder",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:OrderController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:OrderController"],
        beego.ControllerComments{
            Method: "GetListOrder",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:OrderController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:OrderController"],
        beego.ControllerComments{
            Method: "GetListOrderForAdmin",
            Router: "/listAll",
            AllowHTTPMethods: []string{"get"},
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
            Method: "UpDateProduct",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:ProductController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetProduct",
            Router: "/:pid",
            AllowHTTPMethods: []string{"get"},
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

    beego.GlobalControllerRouter["new-order-food/controllers:RateController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:RateController"],
        beego.ControllerComments{
            Method: "SetRate",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:UserController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new-order-food/controllers:UserController"] = append(beego.GlobalControllerRouter["new-order-food/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
