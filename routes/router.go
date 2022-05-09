package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

//路由入口文件


func InitRouter(){

	gin.SetMode(utils.AppMode)
	engine := gin.Default()

	routerV1 := engine.Group("api/v1")
	{
		//用户模块的路由接口
		routerV1.POST("user/add",v1.AddUser)
		routerV1.GET("users",v1.GetUsers)
		routerV1.PUT("user/:id",v1.EditUser)
		routerV1.DELETE("user/:id",v1.DeleteUser)
		//分类模块路由接口

		//文章模块路由接口
	}

	engine.Run(utils.HttpPort)
}
