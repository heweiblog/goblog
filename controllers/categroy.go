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
	login := CheckUser(c.Ctx)
	c.Data["IsLogin"] = login
	c.Data["IsCategory"] = true
	if !login {
		c.TplName = "category.html"
		c.Data["Categorys"] = models.GetAllCategory()
		return
	}

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
	c.TplName = "category_admin.html"
	c.Data["Categorys"] = models.GetAllCategory()
}
