package glogs

import (
	"github.com/astaxie/beego"
)

var Log *logs.BeeLogger

func init() {
	Log = logs.NewLogger(10)
	Log.SetLogger("console", "")
	Log.EnableFuncCallDepth(true)
}
