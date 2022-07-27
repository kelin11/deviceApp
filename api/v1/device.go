package v1

import (
	errmsg "deviceApp/code"
	"deviceApp/model/orm"
	"github.com/gin-gonic/gin"
	"net/http"
)

//查找一个实验室中所有的不可用的设备
func GetAllUnDevice(c *gin.Context) {
	data, code := orm.GetAllUnDevice()
	if code == errmsg.ERROR {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetMessage(code),
		"data":    data,
	})
}
