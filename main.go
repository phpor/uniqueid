package main

import (
	_ "github.com/phpor/uniqueid/routers"
	"github.com/astaxie/beego"
)

func main() {
//	println("test")
	beego.BeeLogger.SetLogger("console", "")
//	beego.Debug(beego.BeeLogger)
	beego.BConfig.Log.AccessLogs = true
	beego.BeeLogger.SetLevel(beego.LevelDebug)
//	beego.Info(beego.BConfig)
	beego.Run()
}

