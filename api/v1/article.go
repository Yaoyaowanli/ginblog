package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)


//AddArt 新增文章
func AddArt (c *gin.Context){
	var data model.Article
	//把文章信息绑定到data
	_ = c.ShouldBindJSON(&data)
	log.Println(data)
	code = model.CreateArt(&data)
	log.Println("20",code)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})
}

//GetCateArt 查询分类下所有文章列表
func GetCateArt(c *gin.Context){
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNo,_ := strconv.Atoi(c.Query("pageno"))
	id,_ := strconv.Atoi(c.Param("id"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNo == 0 {
		pageNo = -1
	}
	data,code := model.GetCateArt(id,pageSize,pageNo)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"data":data,
	})
}

//GetArtInfo 查询单个文章信息
func GetArtInfo(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	if id <0{		//不能出现负数
		code=errmsg.ERROR
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":errmsg.GetErrMsg(code),
		})
	}

	art,code:= model.GetArtInfo(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"data":art,
	})
}

//GetArt 查询所有文章列表
func GetArt (c *gin.Context) {
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNo,_ := strconv.Atoi(c.Query("pageno"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNo == 0 {
		pageNo = -1
	}
	art,code := model.GetArt(pageSize,pageNo)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":art,
		"message":errmsg.GetErrMsg(code),
	})
}

//EditArt 更新文章
func EditArt(c *gin.Context) {
	//拿到url参数id
	id,_ := strconv.Atoi(c.Param("id"))
	if id <0{		//不能出现负数，因为category的id是uint
		code=errmsg.ERROR
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":errmsg.GetErrMsg(code),
		})
	}

	var data model.Article
	_ = c.ShouldBindJSON(&data)

	code = model.EditArt(id,&data)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

//DeleteArt 删除文章
func DeleteArt(c *gin.Context) {
	//拿到url携带的id参数
	id,_ := strconv.Atoi(c.Param("id"))

	//在数据库中找到并删除
	code = model.DeleteArt(id)
	//return
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}