package middleware

import (
	"net/http"
	"new-order-food/models"
	"new-order-food/responses"
	"strings"

	"github.com/astaxie/beego/context"
)

var Token = func(ctx *context.Context) {
	token := ctx.Request.Header["Token"]
	if len(token) < 1 {
		ctx.Output.JSON(responses.UnAuthResponse, true, true)
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return
	}
	str, err := ParseToken(token[0])
	if err != nil {
		ctx.Output.JSON(responses.UnAuthResponse, true, true)
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return
	}
	val := strings.Split(str, ":")
	ctx.Request.Header.Set("id", val[0])
	ctx.Request.Header.Set("type", val[1])
}

func ParseToken(token string) (string, error) {
	val, err := models.Get(token)
	if err != nil {
		return "", err
	}
	return val, nil
}
