package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"log"
	"new-order-food/requests"
	"new-order-food/responses"
	"new-order-food/services"
	"strconv"
)

type CommentController struct {
	beego.Controller
}

//@Title Get List Comment
//@Description Get List Comment
//@Summary Lấy danh sách bình luận
// @Param pid query int true "product id"
//@Success 200 {object} responses.ResponseArray
//@Failure 404 {object} responses.ResponseArray
//@router / [get]
func (this *CommentController) GetListComment() {
	defer this.ServeJSON()
	pid, _ := this.GetInt("pid", -1)
	if pid == -1 {
		log.Println("controllers/comment_controller.go:26 ", errors.New("pid not valid"))
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}

	lc, err := services.GetComments(pid)
	if err != nil {
		log.Println("controllers/comment_controller.go:37 ", err)
		this.Data["json"] = responses.ResponseArray{
			Data:       nil,
			TotalCount: 0,
			Error:      responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseArray{
		Data:       lc,
		TotalCount: len(lc),
		Error:      responses.NewErr(responses.Success),
	}
}

//@Title Create Comment
//@Description Create Comment
//@Summary Tạo mới bình luận
// @Params token header string true "Token"
// @Param data body requests.RequestComment true "comment"
//@Success 200 {object} responses.ResponseBool
//@Failure 404 {object} responses.ResponseBool
//@router / [post]
func (this *CommentController) CreateComment() {
	defer this.ServeJSON()
	idtype := this.Ctx.Request.Header.Get("type")
	if idtype != "user" {
		log.Println("controllers/comment_controller.go:66 , typeid is not user ")
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	uid, _ := strconv.Atoi(this.Ctx.Request.Header.Get("id"))
	com := requests.RequestComment{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &com)
	if err != nil {
		log.Println("controllers/comment_controller.go:77 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	err = services.CreateComment(com, uid)
	if err != nil {
		log.Println("controllers/comment_controller.go:85 ", err)
		this.Data["json"] = responses.ResponseBool{
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.ResponseBool{
		Error: responses.NewErr(responses.Success),
	}
}
