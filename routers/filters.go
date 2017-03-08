package routers

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, checkLogin)
}

func checkLogin(ctx *context.Context) {
	// except login
	if strings.Index(ctx.Request.RequestURI, "/login") != -1 {
		return
	}
	// session
	userData := ctx.Input.Session("userData")
	if userData == nil {
		beego.Info("Redirect to login.")
		ctx.Redirect(302, "/login?uri=" + ctx.Request.RequestURI)
	}
}