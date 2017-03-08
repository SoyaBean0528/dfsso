package controllers

import (
	// "github.com/astaxie/beego"
)

var (
	dashboardLayoutData *LayoutData
)

type DashboardController struct {
	LayoutController
}

func init(){
	dashboardLayoutData = &LayoutData{ PageName:"控制台" }
	// Breadcurmb
	dashboardLayoutData.Breadcrumb = make([]map[string]string, 0, 1)
	// 1
	navi := make(map[string]string)
	navi["Ref"] = "/"
	navi["Name"] = "控制台"
	navi["Icon"] = "fa fa-dashboard"
	dashboardLayoutData.Breadcrumb = append(dashboardLayoutData.Breadcrumb, navi) 
}

func (this *DashboardController) Index() {
	this.LayoutController.SetMainLayout(dashboardLayoutData)
	this.TplName = "dashboard.tpl"
}
