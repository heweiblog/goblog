package controllers

import (
	"github.com/astaxie/beego"
)

type ArchtiveController struct {
	beego.Controller
}

func (c *ArchtiveController) Get() {
	c.Data["IsArchtive"] = true
	c.TplName = "archtive.html"
}
