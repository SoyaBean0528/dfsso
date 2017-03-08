package controllers

import (
	"github.com/astaxie/beego"
)

type LayoutBreadcrumb struct {
	Icon string
	Name string
}

type LayoutData struct {
	PageName string 
	Breadcrumb []LayoutBreadcrumb
}

type LayoutController struct {
	beego.Controller
}

func (this *LayoutController) SetMainLayout(data *LayoutData) {
	this.Data["PageName"] = data.PageName
	this.Layout = "layout_main.tpl"
    // this.LayoutSections = make(map[string]string)
    // this.LayoutSections["PageContent"] = "blogs/html_head.tpl"
}