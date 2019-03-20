package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.GetString("op")
	switch op {
	case "add":
		name := c.GetString("CategoryName")
		err := models.AddCategory(name)
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
