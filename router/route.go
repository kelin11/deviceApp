package router

import (
	v1 "deviceApp/api/v1"
	"deviceApp/middleware"
	"deviceApp/settings"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(settings.AppMode)
	r := gin.Default()
	r.Use(middleware.JWTMiddleWare())
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	//分组路由

	//用户登录
	routerV1 := r.Group("api/v1")
	routerV1.POST("user/login", v1.Login)

	//报修信息路由
	//查看该实验室中所有不可用设备
	routerV1.GET("device/disable", v1.GetAllUnDevice)
	routerV1.POST("device/bug/submit", v1.SubmitUnDevice)

	// 查询故障表数据
	routerV1.GET("device/bug/", v1.GetBugList)
	routerV1.GET("device/bug/detail", v1.GetBugDetail)

	//维修信息
	// 获取所有维修反馈数据
	routerV1.GET("device/solution/", v1.ShowALLSolutionTable)
	// 维修人员上传维修反馈
	routerV1.POST("device/solution/submit", v1.SubmitSolutionTable)

	//将数据库存储solution数据转换为:excel文件
	// Excel
	routerV1.GET("device/bug/export", v1.ExportBugInfo)

	r.Run(settings.HttpPort)

}
