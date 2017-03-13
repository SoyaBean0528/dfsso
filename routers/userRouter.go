package routers

import (
	"dreamfish/dfsso/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/user", &controllers.UserController{}, "get:Index")
    beego.Router("/addUser", &controllers.UserController{}, "post:AddUser")
    beego.Router("/delUser", &controllers.UserController{}, "post:DelUser")
    beego.Router("/editUser", &controllers.UserController{}, "post:EditUser")
}
