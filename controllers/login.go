package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	user := c.GetString("UserName")
	pwd := c.GetString("PassWord")
	login := c.GetString("AutoLogin")

	if beego.AppConfig.String("user") == user && beego.AppConfig.String("pwd") == pwd {
		maxpage := 0
		if login {
			maxage = 1 << 31
		}
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsLogin"] = false
	}
	if login == "on" {
	} else {
	}
	c.Redirect("/", 302)
	return
}
