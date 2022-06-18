package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

//路由入口文件


func InitRouter(){

	gin.SetMode(utils.AppMode)
	engine := gin.New()
	//注册中间件，logger是日志中间件，recovery捕获panic
	engine.Use(middleware.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middleware.Cors())

	//有权限的
	auth := engine.Group("api/v1")
	//路由组中间件
		auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		auth.PUT("user/:id",v1.EditUser)
		auth.DELETE("user/:id",v1.DeleteUser)
		//分类模块路由接口
		auth.POST("category/add",v1.AddCate)
		auth.PUT("category/:id",v1.EditCate)
		auth.DELETE("category/:id",v1.DeleteCate)
		//文章模块路由接口
		auth.POST("article/add",v1.AddArt)
		auth.PUT("article/:id",v1.EditArt)
		auth.DELETE("article/:id",v1.DeleteArt)
		//上传文件
		auth.POST("upload",v1.UpLoad)
	}

	router := engine.Group("api/v1")
	{
		router.POST("user/add",v1.AddUser)
		router.GET("users",v1.GetUsers)
		router.GET("category",v1.GetCate)
		router.GET("article",v1.GetArt)
		router.GET("article/list/:id",v1.GetCateArt)
		router.GET("article/info/:id",v1.GetArtInfo)
		router.POST("login",v1.Login)
	}


	engine.Run(utils.HttpPort)
}
