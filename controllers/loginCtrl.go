package controllers

import (
	"dreamfish/dfsso/models/loginModel"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Index() {
	this.Data["Msg"] = "登录"
	this.Data["Color"] = "black"
	this.Data["Uri"] = this.GetString("uri")
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
		this.Data["Uri"] = uri
		this.Data["Msg"] = err.Error()
		this.Data["Color"] = "red"
		this.TplName = "login.tpl"
		return
	}
	// session
	this.SetSession("userData", user)
	this.Ctx.Redirect(302, uri)
}

func (this *LoginController) Logout() {
	userData := this.GetSession("userData")
	if userData != nil {
		this.DelSession("userData")
	}
	this.Ctx.Redirect(302, "/login")
}
