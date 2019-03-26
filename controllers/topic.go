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
	c.TplName = "topic.html"
	c.Data["IsLogin"] = CheckUser(c.Ctx)
	c.Data["Topics"] = models.GetAllTopic(false)
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
	content := c.GetString("TopicContent")
	if len(content) <= 0 {
		beego.Error("topic content eror len =", len(content))
	}
	id := c.GetString("TopicId")
	if len(id) <= 0 {
		err := models.AddTopic(title, category, content)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.ModTopic(id, title, category, content)
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
