package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"goblog/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	c.Data["IsCategory"] = true

	op := c.GetString("op")

	beego.Error("op:", op)
	switch op {
	case "add":
		name := c.GetString("CategoryName")
		if len(name) <= 0 {
			break
		}
		err := models.AddCategory(name)
		logs.Debug(name, err)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	case "del":
		id := c.GetString("id")
		if len(id) <= 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	}
	c.Data["IsLogin"] = CheckUser(c.Ctx)
	c.Data["Categorys"] = models.GetAllCategory()
	c.TplName = "category.html"
}
