package routers

import (
	"dreamfish/dfsso/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/sub", &controllers.MainController{}, "*:Sub")
}
