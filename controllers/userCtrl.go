package controllers

import (
	// "github.com/astaxie/beego"
	"dreamfish/dfsso/models/userModel"
)

var (
	userLayoutData *LayoutData
)

type UserController struct {
	LayoutController
}

func init(){
	userLayoutData = &LayoutData{ PageName:"用户管理" }
	// Breadcurmb
	userLayoutData.Breadcrumb = make([]map[string]string, 0, 1)
	// 1
	navi := make(map[string]string)
	navi["Ref"] = "/"
	navi["Name"] = "控制台"
	navi["Icon"] = "fa fa-dashboard"
	userLayoutData.Breadcrumb = append(userLayoutData.Breadcrumb, navi) 
	// 2
	navi = make(map[string]string)
	navi["Ref"] = "/user"
	navi["Name"] = "用户管理"
	navi["Icon"] = "glyphicon glyphicon-user"
	userLayoutData.Breadcrumb = append(userLayoutData.Breadcrumb, navi) 
	// Scripts
	userLayoutData.Scripts = make([]string, 0, 1)
	userLayoutData.Scripts = append(userLayoutData.Scripts, "../static/js/user.js")
}

func (this *UserController) Index() {
	// main layout 
	this.LayoutController.SetMainLayout(userLayoutData)
	this.TplName = "user.tpl"
	// user list
	userList := userModel.GetUserList()
	this.Data["UserList"] = userList
}
