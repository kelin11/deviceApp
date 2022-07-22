package router

import (
	v1 "deviceApp/api/v1"
	"deviceApp/settings"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(settings.AppMode)
	r := gin.Default()
	//分组路由

	//用户登录
	routerV1 := r.Group("api/v1")
	routerV1.POST("user/login", v1.Login)

	//保修信息路由

	//维修信息

	//消息推送

}
