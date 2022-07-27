package errmsg

const (
	ERROR   = 500
	SUCCESS = 200
	INVALID_PARAM

	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWD_FAIL      = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_OUT_TIME   = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	//code = 2000... article error
	ERROR_ART_NOT_EXISTED = 2001
	//code = 3000... category error
	ERROR_CATE_USED = 3001
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名称错误",
	ERROR_PASSWD_FAIL:      "用户密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "Token不存在",
	ERROR_TOKEN_OUT_TIME:   "Token已过期",
	ERROR_TOKEN_WRONG:      "token不正确",
	ERROR_TOKEN_TYPE_WRONG: "token格式错误",
	ERROR_ART_NOT_EXISTED:  "文章不存在",
	ERROR_CATE_USED:        "分类名称已经存在",
}

func GetMessage(code int) string {
	return codeMsg[code]
}
