package v1

import (
	errmsg "deviceApp/code"
	"deviceApp/model"
	"deviceApp/model/orm"
	"github.com/gin-gonic/gin"
	"net/http"
)

//保修故障设备
func SubmitUnDevice(c *gin.Context) {
	//拿到上报人的学号
	var bug model.Bug
	_ = c.ShouldBindJSON(&bug)
	uName, _ := c.Get("username")
	uNameStr := uName.(string)
	code := orm.SubmitUnDevice(bug, uNameStr)
	if code == errmsg.ERROR {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetMessage(code),
	})

}
