package router

import (
	"deviceApp/settings"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(settings.AppMode)
	r:=gin.Default()
	//分组路由

	//用户登录
	routerV1 := r.Group("api/user")
	routerV1.POST("user/login",)


	//保修信息路由


	//维修信息

	
	//消息推送

	


	
	

}
