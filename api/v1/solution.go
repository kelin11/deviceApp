package v1

import (
	errmsg "deviceApp/code"
	"deviceApp/model"
	"deviceApp/model/orm"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 维修人员提交维修反馈(存疑问)
func SubmitSolutionTable(c *gin.Context) {
	var solution model.Solution
	_ = c.ShouldBindJSON(&solution)
	code := orm.SubmitSolutionData(solution)
	if code == errmsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetMessage(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetMessage(code),
	})
}

// 获取所有维修反馈数据表
func ShowALLSolutionTable(c *gin.Context) {
	code, data := orm.GetAllSolutionTable()
	if code == errmsg.ERROR {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetMessage(code),
		"data":    data,
	})
}

//将故障表数据转换成excel表单
func TransferExcel(c *gin.Context) {

}
