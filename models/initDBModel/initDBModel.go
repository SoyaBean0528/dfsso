package initDBModel

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"

	"dreamfish/dfsso/models/userModel"
)

func createTables() {
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Println(err)
	}
}

func insertDatas() {
	insertUsers()
}

func insertUsers() {
	user := &userModel.User{ Username:"admin", Password:"12", Creator:"system"}
	userModel.AddUser(user)
}

type DBInitializ struct {

}

func RegisterDB() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpass := beego.AppConfig.String("dbpass")
	dbname := beego.AppConfig.String("dbname")
	beego.Debug(fmt.Sprintf("dbhost:%s dbport:%s dbuser:%s dbpass:%s dbname:%s", dbhost, dbport, dbuser, dbpass, dbname))
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbuser, dbpass, dbhost, dbport, dbname)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
}

func InitDB() {
	beego.Debug("Start initialize database.")
	// register
	RegisterDB()
	// create tables
	createTables()
	// insert datas
	insertDatas()
}
