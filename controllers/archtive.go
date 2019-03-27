package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
)

type ArchtiveController struct {
	beego.Controller
}

func (c *ArchtiveController) Get() {
	c.Data["IsArchtive"] = true
	c.TplName = "archtive.html"
	c.Data["IsLogin"] = CheckUser(c.Ctx)
	c.Data["Topics"], c.Data["TopicCount"], c.Data["ViewCount"], c.Data["ReplyCount"] = models.GetAllTopicByView()
}
