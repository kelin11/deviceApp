package v1

import (
	errmsg "deviceApp/code"
	"deviceApp/model/orm"
	"net/http"

	"github.com/gin-gonic/gin"
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

// 获取某实验室中的所有设备
func GetAllDeviceByLab(c *gin.Context) {
	var lab = c.Query("lab")
	data, code := orm.GetAllDevice(lab)
	if code == errmsg.ERROR {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetMessage(code),
		"data":    data,
	})
}
