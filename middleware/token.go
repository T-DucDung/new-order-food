package middleware

import (
	"net/http"
	"new-order-food/redis"

	"github.com/astaxie/beego/context"
)

var Token = func(ctx *context.Context) {
	token := ctx.Request.Header["Token"]
	if len(token) < 1 {
		//Trả về lỗi token ở đây
		//ctx.Output.JSON(responses.UnAuthResponse, true, true)
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return
	}
	val, err := ParseToken(token[0])
	if err != nil {
		//Trả về lỗi token ở đây
		//ctx.Output.JSON(responses.UnAuthResponse, true, true)
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return
	}

	//Xử lý value
}

func ParseToken(token string) (interface{}, error) {
	val, err := redis.Client.Get(token)
	if err != nil {
		return nil, err
	}
	return val, nil
}
