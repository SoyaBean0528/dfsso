package controllers

import (
	"github.com/astaxie/beego"
	"dreamfish/dfsso/models/loginModel"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.tpl"
}

func (this *LoginController) Post() {
	// get params	
	username := this.GetString("username")
	password := this.GetString("password")
	this.Data["username"] = username
	this.Data["password"] = password
	// log 
	beego.Info("Login Username =", username, "Password =", password)
	// login
	user, err := loginModel.Login(username, password)
	if err != nil {
		this.Data["result"] = err.Error()
		this.TplName = "loginResult.tpl"
		return
	}
	// session
	this.SetSession("uid", user.Id)
	this.Ctx.Redirect(302, "/")
}