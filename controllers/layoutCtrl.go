package controllers

import (
	"dreamfish/dfsso/models/userModel"
	"github.com/astaxie/beego"
)

var (
	mainLayoutSideBar []map[string]string
)

type LayoutData struct {
	Uri string
	PageName   string
	Scripts    []string
	Breadcrumb []map[string]string
}

type LayoutController struct {
	beego.Controller
}

func init() {
	// side bar
	mainLayoutSideBar = make([]map[string]string, 0, 2)
	// 1
	bar := make(map[string]string)
	bar["Name"] = "控制台"
	bar["Href"] = "/"
	bar["Icon"] = "fa fa-dashboard"
	mainLayoutSideBar = append(mainLayoutSideBar, bar)
	// 2
	bar = make(map[string]string)
	bar["Name"] = "用户管理"
	bar["Href"] = "/user"
	bar["Icon"] = "glyphicon glyphicon-user"
	mainLayoutSideBar = append(mainLayoutSideBar, bar)
}

func (this *LayoutController) SetMainLayout(data *LayoutData) {
	// input data
	this.Data["Scripts"] = data.Scripts
	this.Data["PageName"] = data.PageName
	this.Data["Breadcrumb"] = data.Breadcrumb
	this.Data["BreadcrumbEnd"] = len(data.Breadcrumb) - 1
	// base data
	this.Data["SideBar"] = mainLayoutSideBar
	userData := this.GetSession("userData")
	if userData != nil {
		this.Data["UserData"] = userData.(*userModel.User)
	}
	this.Layout = "layout_main.tpl"
	// this.LayoutSections = make(map[string]string)
	// this.LayoutSections["PageContent"] = "blogs/html_head.tpl"
}
