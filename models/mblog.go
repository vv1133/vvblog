package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type BlogPost struct {
	Id       bson.ObjectId `_id`
	Caption  string        `bson:"caption"`
	Slug     string        `bson:"slug"`
	Tags     []string      `bson:"tags"`
	Created  int64         `bson:"created"`
	Markdown string        `bson:"markdown"`
	Html     string        `bson:"html"`
	Cover    string        `bson:"cover"`
	Type     string        `bson:"type"`
}

type BlogTag struct {
	Id      bson.ObjectId `_id`
	Caption string        `bson:"caption"`
	Slug    string        `bson:"slug"`
}

var (
	Session *mgo.Session
	DbPost  *mgo.Collection
	DbTag   *mgo.Collection
	DbConf  *mgo.Collection
	//Option  Conf
)

func (this *BlogPost) Save() error {
	if Has(DbPost, bson.M{"_id": this.Id}) {
		return Update(DbPost, bson.M{"_id": this.Id}, bson.M{"$set": bson.M{"caption": this.Caption, "slug": this.Slug, "tags": this.Tags, "markdown": this.Markdown, "html": this.Html, "cover": this.Cover}})
	}

	this.Id = bson.NewObjectId()
	this.Created = time.Now().Unix()
	return Insert(DbPost, this)
}

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport, _ := beego.AppConfig.Int("dbport")
	dbname := beego.AppConfig.String("dbname")
	dbuser := beego.AppConfig.String("dbuser")
	dbpass := beego.AppConfig.String("dbpass")

	userAndPass := dbuser + ":" + dbpass + "@"
	if dbuser == "" || dbpass == "" {
		userAndPass = ""
	}

	url := "mongodb://" + userAndPass + dbhost + ":" + strconv.Itoa(dbport) + "/" + dbname

	var err error
	Session, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	Session.SetMode(mgo.Monotonic, true)
	DbPost = Session.DB(dbname).C("BlogPost")
	DbTag = Session.DB(dbname).C("BlogTag")
	DbConf = Session.DB(dbname).C("BlogConfig")
}
