package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, checkLogin)
}

func checkLogin(ctx *context.Context) {
	// except login
	if ctx.Request.RequestURI == "/login" {
		return
	}
	// session
	_, ok := ctx.Input.Session("uid").(int64)
	if !ok {
		beego.Info("Redirect")
		ctx.Redirect(302, "/login")
	}
}