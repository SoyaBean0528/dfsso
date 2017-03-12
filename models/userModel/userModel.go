package userModel 

import (
	"errors"
	"strings"
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

var (
	errUserExisted = errors.New("用户已存在！")
	errPasswordNil = errors.New("密码不能为空！")
	errUsernameNil = errors.New("用户名不能为空！")
)

// init
func init() {
	orm.RegisterModel(new(User))
}

// for orm
func (this *User) TableName() string {
	return ("User")
}

/**
 * Method
 */
func checkUser(user *User) (error) {
	if user.Username == "" {
		return errUsernameNil
	}	
	if user.Password == "" {
		return errPasswordNil
	}

	return nil
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

// admin id
func GetAdminID() (int) {
	return 1
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

func AddUser(user *User) (error) {
	// check user
	if err:= checkUser(user) ;err != nil {
		return err
	}
	// encrypt password
	user.Password = EncryptPassword(user.Password)
	// insert
	o := orm.NewOrm()
	_, err := o.Insert(user)
	// insert err
	if err != nil {
		if strings.Index(err.Error(), "1062") != -1 {
			return errUserExisted
		}
	}

	return nil 
}