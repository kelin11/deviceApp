package orm

import (
	errmsg "deviceApp/code"
	"deviceApp/model"
)

//验证用户登录
func CheckLogin(data *model.User) (code int, userName string) {
	var user model.User
	db.Where("username like ?", data.UserName).First(&user)
	//是否存在此用户
	if user.ID == 0 {
		return errmsg.ERROR, ""
	}
	//验证密码
	if data.PassWord != user.PassWord {
		return errmsg.ERROR_PASSWD_FAIL, ""
	}
	return errmsg.SUCCESS, user.UserName
}
