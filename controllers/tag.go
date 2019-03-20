package controllers

import (
	"github.com/astaxie/beego"
)

type TagController struct {
	beego.Controller
}

func (c *TagController) Get() {
	c.Data["IsTag"] = true
	c.TplName = "tag.html"
	c.Data["IsLogin"] = CheckUser(c.Ctx)
}
