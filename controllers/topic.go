package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsTopic"] = true
	c.Data["Topics"] = models.GetAllTopic(false)
	if c.Data["IsLogin"] = CheckUser(c.Ctx); c.Data["IsLogin"] == false {
		c.TplName = "topic.html"
		return
	}

	c.TplName = "topic_admin.html"
	c.Data["IsLogin"] = CheckUser(c.Ctx)
}

func (c *TopicController) Post() {
	c.Data["IsTopic"] = true
	title := c.GetString("TopicTitle")
	if len(title) <= 0 {
		beego.Error("topic title eror len =", len(title))
	}
	category := c.GetString("TopicCategory")
	if len(category) <= 0 {
		beego.Error("topic category eror len =", len(category))
	}
	label := c.GetString("TopicLabel")
	if len(label) <= 0 {
		beego.Error("topic label eror len =", len(label))
	}
	content := c.GetString("TopicContent")
	if len(content) <= 0 {
		beego.Error("topic content eror len =", len(content))
	}
	id := c.GetString("TopicId")
	if len(id) <= 0 {
		err := models.AddTopic(title, category, label, content)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.ModTopic(id, title, category, label, content)
		if err != nil {
			beego.Error(err)
		}
	}
	c.Data["IsLogin"] = CheckUser(c.Ctx)
	c.Redirect("/topic", 301)
}

func (c *TopicController) Add() {
	c.Data["IsTopic"] = true
	c.TplName = "topic_add.html"
	c.Data["IsLogin"] = CheckUser(c.Ctx)
}

func (c *TopicController) Del() {
	c.Data["IsTopic"] = true
	c.Data["IsLogin"] = CheckUser(c.Ctx)
	id := c.Ctx.Input.Params()["0"]
	err := models.DelTopic(id)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 302)
}

func (c *TopicController) Mod() {
	c.Data["IsTopic"] = true
	c.TplName = "topic_mod.html"
	id := c.Ctx.Input.Params()["0"]
	topic, err := models.GetTopic(id)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Labels"] = models.GetLabel(id)
	c.Data["Topic"] = topic
	c.Data["TopicId"] = id
}

func (c *TopicController) View() {
	c.Data["IsTopic"] = true
	c.Data["IsLogin"] = CheckUser(c.Ctx)
	c.TplName = "topic_view.html"
	id := c.Ctx.Input.Params()["0"]
	topic, err := models.GetTopic(id)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Topic"] = topic
	c.Data["TopicId"] = id
	c.Data["Comments"] = models.GetAllComment(id, true)
}
