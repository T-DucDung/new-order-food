package test

import (
	"log"
	"new-order-food/models"
	_ "new-order-food/routers"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestExitsToken(t *testing.T) {
	a, err := models.Exist("0109abe657cb0cc23260bf079ad4603c11fb1bae5efbd8216bbd0ca303447c59")
	log.Println("----------------",a, err)
}
