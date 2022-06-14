package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)


//UpLoad 拿到提交的file文件，调用model.UploadFile上传
func UpLoad (c *gin.Context){
	file,fileHeader,_ := c.Request.FormFile("file")

	url,code := model.UploadFile(file,fileHeader.Size)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
		"url" : url,
	})
}
