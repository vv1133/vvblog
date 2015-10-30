package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/vv1133/vvblog/common"
	"github.com/vv1133/vvblog/models"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type ArticlesController struct {
	beego.Controller
}

func (c *ArticlesController) New() {
	c.TplNames = "admin/new.html"
}

func (c *ArticlesController) Edit() {
	ids := c.Ctx.Input.Param(":id")
	if ids == "" {
		c.Abort("404")
	}

	id := bson.ObjectIdHex(ids)

	var scpost models.BlogPost

	models.GetOneById(models.DbPost, id, &scpost)

	c.Data["Id"] = id.Hex()
	c.Data["Caption"] = scpost.Caption
	c.Data["Slug"] = scpost.Slug

	if scpost.Type == "page" {
		c.Data["IsPage"] = true
	} else {
		c.Data["IsPost"] = true
	}

	c.Data["Markdown"] = scpost.Markdown
	c.Data["Cover"] = scpost.Cover
	c.Data["Tags"] = scpost.Tags

	c.TplNames = "admin/new.html"
}

func (c *ArticlesController) Del() {
	ids := c.Ctx.Input.Param(":id")
	if ids == "" {
		c.Data["json"] = map[string]interface{}{"error": "ID Error!"}
	} else {
		id := bson.ObjectIdHex(ids)
		err := models.Delete(models.DbPost, bson.M{"_id": id})
		if err != nil {
			c.Data["json"] = map[string]interface{}{"error": err.Error()}
		} else {
			c.Data["json"] = map[string]interface{}{"error": "0"}
		}
	}

	c.ServeJson()
}

func (c *ArticlesController) Update() {
	ids := c.GetString("id")
	id := bson.NewObjectId()
	if ids != "" {
		id = bson.ObjectIdHex(ids)
	}

	caption := c.GetString("caption")
	slug := c.GetString("slug")
	atype := c.GetString("type")
	markdown := c.GetString("editor-markdown-doc")
	html := c.GetString("html")
	cover := c.GetString("cover")
	tag := c.GetString("tag")
	splits := strings.Split(tag, ",")

	var tags []string

	if len(splits) > 0 && splits[0] != "" {
		for _, v := range splits {
			tags = append(tags, strings.TrimSpace(v))
			s := common.GetSlug(v, false)
			models.Tag(strings.TrimSpace(v), strings.TrimSpace(s))
		}
	} else {
		tags = splits
	}

	fmt.Println("ids:", ids)
	fmt.Println("caption:", caption)
	fmt.Println("slug:", slug)
	fmt.Println("atype:", atype)
	fmt.Println("markdown:", markdown)
	fmt.Println("html:", html)
	fmt.Println("cover:", cover)
	fmt.Println("tag:", tag)
	fmt.Println("splits:", splits)

	scpost := &models.BlogPost{
		Id:       id,
		Caption:  caption,
		Slug:     slug,
		Tags:     tags,
		Markdown: markdown,
		Html:     html,
		Cover:    cover,
		Type:     atype,
	}

	err := scpost.Save()
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/admin", 302)
}
