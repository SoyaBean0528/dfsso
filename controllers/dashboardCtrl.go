package controllers

import (
	// "github.com/astaxie/beego"
)

var (
	layoutData *LayoutData
)

type DashboardController struct {
	LayoutController
}

func init(){
	layoutData = &LayoutData{ PageName:"Dashboard" }
}

func (this *DashboardController) Index() {
	this.LayoutController.SetMainLayout(layoutData)
	this.TplName = "dashboard.tpl"
}
