package router

import (
	"deviceApp/api/v1"
	"deviceApp/middleware"
	"deviceApp/settings"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(settings.AppMode)
	r := gin.Default()
	r.Use(middleware.JWTMiddleWare())
	//分组路由

	//用户登录
	routerV1 := r.Group("api/v1")
	routerV1.POST("user/login", v1.Login)

	//报修信息路由
	//查看该实验室中所有不可用设备
	routerV1.GET("device/", v1.GetAllUnDevice)
	routerV1.POST("device/submit", v1.SubmitUnDevice)

	// 根据实验室Lab查询设备
	routerV1.GET("device/", v1.GetAllDeviceByLab)

	// 查询故障表数据
	routerV1.GET("bug/", v1.ShowBugTable)
	// 根据实验室Lab查询故障数据表
	routerV1.GET("bug/", v1.ShowBugTableByLab)

	//维修信息
	// 获取所有维修反馈数据
	routerV1.GET("device/solution/getData", v1.ShowALLSolutionTable)
	// 维修人员上传维修反馈
	routerV1.POST("device/solution/submit", v1.SubmitSolutionTable)


	r.Run(settings.HttpPort)

}
