package v1

import (
	errmsg "deviceApp/code"
	"deviceApp/model"
	"deviceApp/model/orm"
	"github.com/gin-gonic/gin"
	"net/http"
)

//维修人员的账号数据库里面有吗?
func Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	code := orm.CheckLogin(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": errmsg.GetMessage(code),
		"code":    code,
	})
}
