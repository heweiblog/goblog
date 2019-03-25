package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	c.Data["IsCategory"] = true
	if c.Data["IsLogin"] = CheckUser(c.Ctx); c.Data["IsLogin"] == false {
		c.Data["Categorys"] = models.GetAllCategory()
		c.TplName = "category.html"
		return
	}

	name := c.GetString("CategoryName")
	if len(name) > 0 {
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
	}

	c.Data["Categorys"] = models.GetAllCategory()
	c.TplName = "category_admin.html"
}

func (c *CategoryController) Post() {
	c.Data["IsCategory"] = true
	if c.Data["IsLogin"] = CheckUser(c.Ctx); c.Data["IsLogin"] == false {
		c.Data["Categorys"] = models.GetAllCategory()
		c.TplName = "category.html"
		return
	}

	name := c.GetString("CategoryName")
	if len(name) <= 0 {
		c.Redirect("/category", 301)
		return
	}
	id := c.GetString("CategoryId")
	if len(id) <= 0 {
		c.Redirect("/category", 301)
		return
	}
	beego.Error(id, name)
	err := models.ModCategory(id, name)
	if err != nil {
		beego.Error(err)
	}

	c.Data["Categorys"] = models.GetAllCategory()
	c.Redirect("/category", 301)
	return
}

func (c *CategoryController) Mod() {
	c.TplName = "category_mod.html"
	id := c.Ctx.Input.Params()["0"]
	category, err := models.GetCategory(id)
	if err != nil {
		beego.Error(err)
		c.Redirect("/category", 301)
		return
	}
	c.Data["Category"] = category
}

func (c *CategoryController) Del() {
	id := c.Ctx.Input.Params()["0"]
	beego.Error(id)
	err := models.DelCategory(id)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/category", 301)
	return
}
