package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
)

type TagController struct {
	beego.Controller
}

func (c *TagController) Get() {
	c.Data["IsTag"] = true
	c.TplName = "tag.html"
	c.Data["IsLogin"] = CheckUser(c.Ctx)
	c.Data["Labels"] = models.GetAllLabel(false)
}

func (c *TagController) View() {
	c.Data["IsTag"] = true
	c.Data["IsLogin"] = CheckUser(c.Ctx)
	c.TplName = "tag_view.html"
	id := c.Ctx.Input.Params()["0"]
	tag, err := models.GetTag(id)
	if err != nil {
		beego.Error(err)
		c.Redirect("/tag", 302)
		return
	}
	c.Data["Label"] = tag
	c.Data["LabelId"] = id
	topics, err := models.GetAllTopicByLabel(id)
	if err != nil {
		beego.Error(err)
		c.Redirect("/tag", 302)
		return
	}
	c.Data["Topics"] = topics
}
