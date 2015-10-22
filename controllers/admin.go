package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/vv1133/vvblog/models"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	var scposts []models.BlogPost

	count := models.Count(models.DbPost, nil)
	page := pagination.NewPaginator(c.Ctx.Request, 10, count)
	c.Data["paginator"] = page

	models.GetDataByQuery(models.DbPost, page.Offset(), 10, "-created", nil, &scposts)
	c.Data["Lists"] = scposts

	c.TplNames = "admin/admin.html"
}

func (c *AdminController) Login() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "admin/login.html"
}

func (c *AdminController) Logout() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "admin/login.html"
}
