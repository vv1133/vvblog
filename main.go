package main

import (
	"github.com/astaxie/beego"
	"github.com/vv1133/vvblog/common"
	_ "github.com/vv1133/vvblog/routers"
)

func main() {
	beego.AddFuncMap("LoadTimes", common.LoadTimes)
	beego.AddFuncMap("GetTagSlug", common.GetTagSlug)
	beego.Run()
}
