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
	//维修信息

	r.Run(settings.HttpPort)

}
