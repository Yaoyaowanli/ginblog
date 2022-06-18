package model

import (
	"context"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"mime/multipart"
)

var AccessKey = utils.AccessKey  //七牛ak密钥
var SecretKey = utils.SecretKey  //七牛sk密钥
var Bucket = utils.Bucket    //七牛空间名
var ImgUrl = utils.QiniuServer


//UploadFile  上传文件
func UploadFile (file multipart.File,fileSize int64)(string,int){
	PutPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey,SecretKey)
	upToken := PutPolicy.UploadToken(mac)
	//cfg 配置上传，只配置七牛机房所在位置，其余不做配置（要钱）
	cfg := storage.Config{
		Zone: &storage.ZoneHuanan,
		UseHTTPS: false,
		UseCdnDomains: false,
	}
	//putExtra 表单上传的额外可选项,这里不做任何配置
	putExtra := storage.PutExtra{}
	//formUploader 构建一个表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	//ret 上传回复内容
	ret := storage.PutRet{}
	//PutWithoutKey 用来以表单方式上传一个文件。
	err := formUploader.PutWithoutKey(context.Background(),&ret,upToken,file,fileSize,&putExtra)
	if err != nil {
		log.Println(err)
		return "",errmsg.ERROR
	}
	//上传成功
	url := ImgUrl+ret.Key
	return url,errmsg.SUCCESS
}