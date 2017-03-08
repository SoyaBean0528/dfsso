package routers

import (
	"dreamfish/dfsso/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/login", &controllers.LoginController{}, "get:Index;post:Login")
    beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
}
