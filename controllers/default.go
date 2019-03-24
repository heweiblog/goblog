package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"goblog/models"
)

type MainController struct {
	beego.Controller
}

// 检查用户是否登陆
func CheckUser(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("user")
	if err != nil {
		return false
	}
	user := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	return beego.AppConfig.String("user") == user && beego.AppConfig.String("pwd") == pwd
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "index.html"
	c.Data["IsLogin"] = CheckUser(c.Ctx)
	c.Data["Topics"] = models.GetAllTopic(true)
}
