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
	var uName string
	var bug model.Bug
	_ = c.ShouldBindJSON(&bug)
	uName = c.Param(uName)
	code := orm.SubmitUnDevice(bug, uName)
	if code == errmsg.ERROR {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetMessage(code),
	})

}


// 故障表数据展示接口--无条件查询
func ShowBugTable(c *gin.Context) {
	data, code := orm.GetAllBugTable()
	if code == errmsg.ERROR{
		c.JSON(http.StatusOK, gin.H{
			"code": 	code,
			"message": 	errmsg.GetMessage(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 	code,
		"message": 	errmsg.GetMessage(code),
		"data": 	data,
	})
}

// 根据实验室Lab故障表数据展示接口--条件查询
func ShowBugTableByLab(c *gin.Context)  {
	lab := c.Query("lab")
	code, data := orm.GetBugTableByLab(lab)
	if code != errmsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"code": 	code,
			"message": 	errmsg.GetMessage(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 	code,
		"message": 	errmsg.GetMessage(code),
		"data": 	data,
	})
}