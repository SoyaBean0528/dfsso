package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "main.tpl"
}

func (this *MainController) Sub() {
	this.TplName = "sub.tpl"	
}