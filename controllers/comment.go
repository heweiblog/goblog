package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
)

type CommentController struct {
	beego.Controller
}

func (c *CommentController) Post() {
	c.Data["IsComment"] = true
	name := c.GetString("NickName")
	if len(name) <= 0 {
		beego.Error("nickname eror len =", len(name))
	}
	email := c.GetString("Email")
	if len(email) <= 0 {
		beego.Error("email eror len =", len(email))
	}
	content := c.GetString("Content")
	if len(content) <= 0 {
		beego.Error("comment content eror len =", len(content))
		c.Redirect("/", 301)
		return
	}
	id := c.GetString("TopicId")
	if len(id) <= 0 {
		beego.Error("comment id eror len =", len(id))
		c.Redirect("/", 301)
		return
	}
	err := models.AddComment(id, name, email, content)
	if err != nil {
		beego.Error(err)
	}
	c.Data["IsLogin"] = CheckUser(c.Ctx)
	c.Redirect("/topic/view/"+id, 301)
}
