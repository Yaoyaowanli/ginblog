package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//AddCate 添加类别
func AddCate(c *gin.Context)  {
	//1.拿到用户名
	var data model.Category
	//取出前端传来的数据绑定到 data结构体
	_ = c.ShouldBindJSON(&data)
	//2、在数据库中检查有无同名，如果有错误，抛出
	code = model.CheckCategory(data.Id)
	if code == errmsg.SUCCESS{		//有重名
		code=errmsg.ERROR_CATENAME_USED
	}else if code == errmsg.ERROR_CATE_NOT_EXIST{
		model.CreateCate(&data)
		code = errmsg.SUCCESS
	}

	//3.返回http响应
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})
}

//查询单个用户


//GetCate 查询类别列表
func GetCate(c *gin.Context){
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNo,_ := strconv.Atoi(c.Query("pageno"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNo == 0 {
		pageNo = -1
	}
	cate := model.GetCate(pageSize,pageNo)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":cate,
		"message":errmsg.GetErrMsg(code),
	})
}

//EditCate 编辑类别
func EditCate(c *gin.Context){
	//拿到url参数id
	id,_ := strconv.Atoi(c.Param("id"))
	if id <0{		//不能出现负数，因为category的id是uint
		code=errmsg.ERROR
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":errmsg.GetErrMsg(code),
		})
	}
	//新建模型
	var data model.Category
	//解析数据到data
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Id)
	if code == errmsg.SUCCESS{
		model.EditCate(id,&data)
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":errmsg.GetErrMsg(code),
		})
	}else if code == errmsg.ERROR_CATE_NOT_EXIST {
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":errmsg.GetErrMsg(code),
		})
	}
}

//DeleteCate 删除类别
func DeleteCate(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	if id < 0 {
		code = errmsg.ERROR_CATE_WRONG
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":errmsg.GetErrMsg(code),
		})
		return
	}
	cate := model.FindByCateId(id)
	if cate != nil{
		code = model.DeleteCate(id)
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":errmsg.GetErrMsg(code),
		})
		return
	}else {
		code = errmsg.ERROR
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":errmsg.GetErrMsg(code),
		})
	}

}
