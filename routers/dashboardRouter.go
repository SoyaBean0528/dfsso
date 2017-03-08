package routers

import (
	"dreamfish/dfsso/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.DashboardController{}, "get:Index")
}
