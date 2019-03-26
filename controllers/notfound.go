package controllers

import (
	"github.com/astaxie/beego"
)

type NotFoundController struct {
	beego.Controller
}

func (c *NotFoundController) Get() {
	c.TplName = "404.html"
}
