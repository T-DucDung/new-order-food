package test

import (
	"new-order-food/models"
	_ "new-order-food/routers"
	"testing"
)

func init() {
	models.InitConnectDataBase()
	models.InitRedisClient()
}

func TestGetData(t *testing.T) {

}
