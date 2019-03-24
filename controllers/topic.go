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
	c.Data["Topics"] = models.GetAllTopic()
}

func (c *TopicController) Post() {
	c.Data["IsTopic"] = true
	title := c.GetString("TopicTitle")
	if len(title) <= 0 {
		beego.Error("topic title eror len =", len(title))
	}
	content := c.GetString("TopicContent")
	if len(content) <= 0 {
		beego.Error("topic content eror len =", len(content))
	}
	err := models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}
	c.Data["IsLogin"] = CheckUser(c.Ctx)
	c.Redirect("/topic", 301)
}

func (c *TopicController) Add() {
	c.Data["IsTopic"] = true
	c.TplName = "topic_add.html"
	c.Data["IsLogin"] = CheckUser(c.Ctx)
}
