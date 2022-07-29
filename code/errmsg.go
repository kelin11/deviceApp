package errmsg

const (
	ERROR   = 500
	SUCCESS = 200
	INVALID_PARAM

	ERROR_USERNAME_USED = 1001
	ERROR_PASSWD_FAIL   = 1002
)

var codeMsg = map[int]string{
	SUCCESS:             "OK",
	ERROR:               "FAIL",
	ERROR_USERNAME_USED: "用户名称错误",
	ERROR_PASSWD_FAIL:   "用户密码错误",
}

func GetMessage(code int) string {
	return codeMsg[code]
}
