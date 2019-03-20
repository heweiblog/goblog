package controllers

import (
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsTopic"] = true
	c.TplName = "topic.html"
	c.Data["IsLogin"] = CheckUser(c.Ctx)
}
