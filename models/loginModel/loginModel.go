package loginModel

import (
	"errors"
	"dreamfish/dfsso/models/userModel"
)

func Login(username string, password string) (user *userModel.User, err error) {
	// get by name
	user = userModel.ByName(username)		
	if user.Id == 0 {
		return nil, errors.New("用户不存在!")
	}
	// check password
	if !userModel.CheckPassword(user, password) {
		return nil, errors.New("密码错误!")
	} 

	return user, nil 
}