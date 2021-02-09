package main

import (
	"new-order-food/models"
	"new-order-food/redis"
	_ "new-order-food/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.BConfig.WebConfig.AutoRender = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/v1/swagger"] = "swagger"

	models.InitConnectDataBase()
	redis.InitRedisClient()

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "content-type", "Content-Type", "Auth"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// beego.InsertFilter("/v1/statisticstore/*", beego.BeforeRouter, middlewares.Jwt)
	// beego.InsertFilter("/v1/statistic/*", beego.BeforeRouter, middlewares.Jwt)
	// beego.InsertFilter("/v1/listorder/*", beego.BeforeRouter, middlewares.Jwt)
	// beego.InsertFilter("/v1/user/*", beego.BeforeRouter, middlewares.Jwt)
	// beego.InsertFilter("/v1/order/*", beego.BeforeRouter, middlewares.Jwt)
	// beego.InsertFilter("/v1/store/comment", beego.BeforeRouter, middlewares.Jwt)
	// beego.InsertFilter("/v1/store/rate", beego.BeforeRouter, middlewares.Jwt)

	// beego.InsertFilter("*", beego.BeforeRouter)

	beego.Run()
}
