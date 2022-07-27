package v1

import (
	errmsg "deviceApp/code"
	"deviceApp/middleware"
	"deviceApp/model"
	"deviceApp/model/orm"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Username string

//维修人员的账号数据库里面有吗?
func Login(ctx *gin.Context) {
	var user model.User
	_ = ctx.ShouldBindJSON(&user)
	code, username := orm.CheckLogin(&user)
	if code == errmsg.ERROR {
		ctx.Abort()
	}
	Username = username
	ctx.JSON(http.StatusOK, gin.H{
		"message": errmsg.GetMessage(code),
		"code":    code,
	})
	middleware.CreateToken(ctx)
}
