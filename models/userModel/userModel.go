package userModel 

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
)

//用户表
type User struct {
	Id       int64  `orm:"pk;auto"`
	Username string `orm:"index;unique;size(20)"`
	Password string `orm:"size(32)"`
	Creator  string `orm:"size(20)"`
}

// init
func init() {
	orm.RegisterModel(new(User))
}

// for orm
func (this *User) TableName() string {
	return ("User")
}

/**
 * Public Interface
 */

// encrypt the password by md5
func EncryptPassword(password string) string {
	m := md5.New()
    m.Write([]byte(password))

	return hex.EncodeToString(m.Sum(nil))
}

// check password
func CheckPassword(user *User, password string) bool {
	return user.Password == EncryptPassword(password)	
}

// get user by username
func ByName(username string) (user *User) {
	user = &User{Username:username}

	o := orm.NewOrm()
	o.Read(user, "Username")

	return user
}

func GetUserList() (users []*User) {
	user := new(User)

	o := orm.NewOrm()
	o.QueryTable(user).All(&users)

	return users
}