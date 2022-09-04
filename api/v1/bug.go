package v1

import (
	errmsg "deviceApp/code"
	"deviceApp/model"
	"deviceApp/model/orm"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//保修故障设备
func SubmitUnDevice(c *gin.Context) {
	//拿到上报人的学号
	var bug model.Bug
	_ = c.ShouldBindJSON(&bug)

	uName, _ := c.Get("username")
	fmt.Println(uName)
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

// 故障表数据展示接口--无条件查询
func GetBugList(c *gin.Context) {
	data, code := orm.GetBugList()
	if code == errmsg.ERROR {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetMessage(code),
		"data":    data,
	})
}

//展示故障详情 需要bugId
func GetBugDetail(c *gin.Context) {
	bid := c.Query("id")
	// 将string类型转换为int类型
	bugId, _ := strconv.Atoi(bid)
	data, code := orm.GetBugDetail(bugId)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetMessage(code),
		"data": data,
	})
}

// ExportBugInfo
func ExportBugInfo(c *gin.Context) {
	excel := orm.ExportBugExcel()
	c.Header("response-type", "blob")
	c.Data(http.StatusOK, "application/vnd.ms-excel", excel.Bytes())
}
