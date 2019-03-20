package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	if c.GetString("exist") == "true" {
		c.Ctx.SetCookie("user", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Redirect("/", 302)
		return
	}
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	//c.Ctx.WriteString(fmt.Sprint(c.Input()))
	user := c.GetString("UserName")
	pwd := c.GetString("PassWord")
	autologin := c.GetString("AutoLogin")

	if beego.AppConfig.String("user") == user && beego.AppConfig.String("pwd") == pwd {
		maxpage := 0
		if autologin == "on" {
			maxpage = 1 << 31
		}
		c.Ctx.SetCookie("user", user, maxpage, "/")
		c.Ctx.SetCookie("pwd", pwd, maxpage, "/")
	}
	c.Redirect("/", 302)
	return
}
