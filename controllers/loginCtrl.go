package controllers

import (
	"github.com/astaxie/beego"
	"dreamfish/dfsso/models/loginModel"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Index() {
	this.Data["msg"] = "登录"
	this.Data["color"] = "black"
	this.Data["uri"] = this.GetString("uri") 
	this.TplName = "login.tpl"
}

func (this *LoginController) Login() {
	// get params	
	uri := this.GetString("uri")
	username := this.GetString("username")
	password := this.GetString("password")
	this.Data["username"] = username
	this.Data["password"] = password
	// log 
	beego.Info("Login Username =", username, "Password =", password, "Uri = ", uri)
	// login
	user, err := loginModel.Login(username, password)
	if err != nil {
		this.Data["msg"] = err.Error()
		this.Data["color"] = "red"
		this.TplName = "login.tpl"
		return
	}
	// session
	this.SetSession("uid", user.Id)
	this.Ctx.Redirect(302, uri)
}