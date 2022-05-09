package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

//UserExist 查询用户是否存在
func UserExist(c *gin.Context){

}

//AddUser 添加用户
func AddUser(c *gin.Context)  {
	//1.拿到用户名
	var data model.User
	//取出前端传来的数据绑定到 data结构体
	_ = c.ShouldBindJSON(&data)
	//2、在数据库中检查有无同名，如果有错误，抛出
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS{		//没有重名，进行添加
		model.CreateUser(&data)
	}else if code == errmsg.ERROR_USERNAME_USED{
		code = errmsg.ERROR_USERNAME_USED
	}

	//3.返回http响应
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})
}

//查询单个用户


//GetUsers 查询用户列表
func GetUsers(c *gin.Context){
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNo,_ := strconv.Atoi(c.Query("pageno"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNo == 0 {
		pageNo = -1
	}
	users := model.GetUsers(pageSize,pageNo)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":users,
		"message":errmsg.GetErrMsg(code),
	})
}

//EditUser 编辑用户
func EditUser(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS{
		model.EditUser(id,&data)
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":errmsg.GetErrMsg(code),
		})
	}else if code == errmsg.ERROR_USERNAME_USED {
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":errmsg.GetErrMsg(code),
		})
	}
}

//DeleteUser 删除用户
func DeleteUser(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	user := model.FindByID(id)
	if user != nil{
		code = model.DeleteUser(id)
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":errmsg.GetErrMsg(code),
		})
		return
	}
	code = errmsg.ERROR
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}