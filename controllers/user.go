package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/vv1133/vvblog/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Index() {
	var scposts []models.BlogPost

	count := models.Count(models.DbPost, nil)
	page := pagination.NewPaginator(c.Ctx.Request, 10, count)
	c.Data["paginator"] = page

	models.GetDataByQuery(models.DbPost, page.Offset(), 10, "-created", nil, &scposts)
	c.Data["Lists"] = scposts

	c.Data["SiteName"] = beego.AppConfig.String("sitename")
	c.Data["SubTitle"] = beego.AppConfig.String("subtitle")

	c.TplNames = "user/index.html"
}

func (c *UserController) View() {
	slug := c.Ctx.Input.Param(":slug")
	if slug == "" {
		c.Abort("404")
	}

	var scpost models.BlogPost

	models.GetOneByQuery(models.DbPost, bson.M{"slug": slug}, &scpost)

	if scpost.Id.Hex() == "" {
		c.Abort("404")
	}

	c.Data["Caption"] = scpost.Caption
	c.Data["Type"] = scpost.Type
	c.Data["Created"] = scpost.Created
	c.Data["Tags"] = scpost.Tags
	c.Data["Html"] = scpost.Html
	c.Data["Id"] = scpost.Id.Hex()
	c.Data["Slug"] = scpost.Slug
	c.Data["StartTime"] = time.Now()
	c.TplNames = "user/view.html"
}

func (this *UserController) TagList() {
	tag := this.Ctx.Input.Param(":tag")
	if tag == "" {
		this.Abort("404")
	}

	var sctag models.BlogTag

	models.GetOneByQuery(models.DbTag, bson.M{"slug": tag}, &sctag)

	if sctag.Id.Hex() == "" {
		this.Data["NotSearch"] = true
		this.TplNames = "home/index.html"
		return
	}

	tag = sctag.Caption

	//      this.Data["SiteName"] = tag
	//      this.Data["SubTitle"] = models.Option.SiteName
	//      this.Data["Keywords"] = models.Option.Keywords
	//      this.Data["Description"] = models.Option.Description

	var scposts []models.BlogPost

	count := models.Count(models.DbPost, bson.M{"type": "post", "tags": tag})

	page := pagination.NewPaginator(this.Ctx.Request, 10, count)

	this.Data["paginator"] = page

	models.GetDataByQuery(models.DbPost, page.Offset(), 10, "-created", bson.M{"type": "post", "tags": tag}, &scposts)
	this.Data["Lists"] = scposts

	if len(scposts) <= 0 {
		this.Data["NotSearch"] = true
	}

	this.TplNames = "user/index.html"
}
