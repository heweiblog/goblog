package routers

import (
	"github.com/astaxie/beego"
	"goblog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/archtive", &controllers.ArchtiveController{})
	beego.Router("/tag", &controllers.TagController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/comment", &controllers.CommentController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.AutoRouter(&controllers.CategoryController{})
}
