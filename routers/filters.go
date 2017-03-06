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
	_, ok := ctx.Input.Session("uid").(int64)
	if !ok {
		beego.Info("Redirect to login.")
		ctx.Redirect(302, "/login?uri=" + ctx.Request.RequestURI)
	}
}