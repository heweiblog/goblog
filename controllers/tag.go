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
	c.Data["Labels"] = models.GetAllLabel(false)
	if c.Data["IsLogin"] = CheckUser(c.Ctx); c.Data["IsLogin"] == false {
		c.TplName = "tag.html"
		return
	}
	c.TplName = "tag_admin.html"
}

func (c *TagController) Del() {
	c.Data["IsTag"] = true
	c.Data["IsLogin"] = CheckUser(c.Ctx)
	id := c.Ctx.Input.Params()["0"]
	err := models.DelLabel(id)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/tag", 302)
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
