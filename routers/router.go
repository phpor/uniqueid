package routers

import (
	"github.com/astaxie/beego"
	"github.com/phpor/uniqueid/controllers/uniqueid"
	"github.com/phpor/uniqueid/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
//	beego.Include(&uniqueid.UniqueidController{})
	beego.Router("/api", &uniqueid.UniqueidController{}, "get:Get")
}
