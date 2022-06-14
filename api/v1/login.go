package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login (c *gin.Context) {
	var data model.User
	var token string
	//获取绑定结构体
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("ShouldBindJSON(&data) failed,err:",err)
	}
	//登陆验证
	code = model.CheckLogin(data.Username,data.Password)
	if code == errmsg.SUCCESS{
		//如果验证成功，就生成一个token返回
		token,code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"token":token,
	})
}
