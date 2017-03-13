package controllers

import (
	"github.com/astaxie/beego"
	"dreamfish/dfsso/models/userModel"
)

var (
	userLayoutData *LayoutData
)

type UserController struct {
	LayoutController
}

func init() {
	userLayoutData = &LayoutData{PageName: "用户管理", Uri:"/user"}
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
	// err flash
	flash := beego.ReadFromRequest(&this.Controller)
	if err := flash.Data["error"]; err != "" {
		msg := flash.Data["notice"]
		this.Data[msg + "Msg"] = err 
    }
    this.Data["AdminID"] = userModel.GetAdminID() 
	this.Data["UserList"] = userList
}

func (this *UserController) AddUser() {
	// session
	userData := this.GetSession("userData").(*userModel.User)
	// user 
	user := new(userModel.User)
	user.Creator  = userData.Username
	user.Username = this.GetString("username")	
	// log
	beego.Info("Add User Username =", user.Username, "Creator =", user.Creator)
	// add
	err := userModel.AddNewUser(user)
	if err != nil {
		flash := beego.NewFlash()
		flash.Notice("AddUser")
		flash.Error(err.Error())
		flash.Store(&this.Controller)
	}

	this.Redirect("/user", 302)
}

func (this *UserController) DelUser() {
	// var
	id, _ := this.GetInt("userid")
	userData := this.GetSession("userData").(*userModel.User)
	// uer 
	user := &userModel.User{ Id:int64(id) }
	// log
	beego.Info("Del User UserID =", user.Id)
	// del
	err := userModel.DelUser(user, userData)
	if err != nil {
		flash := beego.NewFlash()
		flash.Notice("DelUser")
		flash.Error(err.Error())
		flash.Store(&this.Controller)
	}
	
	this.Redirect("/user", 302)
}

func (this *UserController) EditUser() {
	// var
	id, _ := this.GetInt("userid")
	// user
	user := new(userModel.User)
	user.Id = int64(id)
	user.Username = this.GetString("username")
	user.Password = this.GetString("password")
	beego.Info("UpdateUser UserName = ", user.Username)
	// update
	err := userModel.UpdateUser(user)
	if err != nil {
		flash := beego.NewFlash()
		flash.Notice("UpdateUser")
		flash.Error(err.Error())
		flash.Store(&this.Controller)
	}

	this.Redirect("/user", 302)
}