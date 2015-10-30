package routers

import (
	"github.com/astaxie/beego"
	"github.com/vv1133/vvblog/controllers"
)

func init() {
	beego.SetStaticPath("/static", "static")

	beego.ErrorController(&controllers.ErrorController{})

	beego.Router("/", &controllers.UserController{}, "get:Index")
	beego.Router("/:slug", &controllers.UserController{}, "get:View")
	beego.Router("/tag/:tag", &controllers.UserController{}, "get:TagList")

	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/admin/login", &controllers.AdminController{}, "get:Login")
	beego.Router("/admin/logout", &controllers.AdminController{}, "get:Logout")

	beego.Router("/admin/new", &controllers.ArticlesController{}, "get:New")
	beego.Router("/admin/edit/:id", &controllers.ArticlesController{}, "get:Edit")
	beego.Router("/admin/del/:id", &controllers.ArticlesController{}, "get:Del")
	beego.Router("/admin/update", &controllers.ArticlesController{}, "post:Update")

	beego.Router("/admin/setting", &controllers.SettingController{}, "get:Get")
}
