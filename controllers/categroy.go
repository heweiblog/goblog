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
	op := c.GetString("op")

	logs.Debug("op:", op)
	switch op {
	case "add":
		name := c.GetString("CategoryName")
		err := models.AddCategory(name)
		logs.Debug(name, err)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	case "del":
	}
	c.Data["IsCategory"] = true
	c.TplName = "category.html"
	c.Data["IsLogin"] = CheckUser(c.Ctx)
	c.Data["Category"] = models.GetAllCategory()
}
