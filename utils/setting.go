package utils

//读取config文件，取出参数

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var ( // server
	AppMode string
	HttpPort string
	JwtKey string
	//  database
	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassword string
	DbName string
	//QiNiu
	AccessKey string
	SecretKey string
	Bucket string
	QiniuServer string
)

//init 读取ini文件
func init() {
	file,err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("read config failed,err:",err)
	}
	LoadServer(file)
	LoadDb(file)
	LoadQiNiu(file)
}

//LoadServer 读取config中server参数
func LoadServer(file *ini.File){
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")

}

//LoadDb 读取config中dataBase参数
func LoadDb(file *ini.File){
	Db = file.Section("dataBase").Key("Db").MustString("mysql")
	DbHost = file.Section("dataBase").Key("DbHost").MustString("localhost")
	DbPort = file.Section("dataBase").Key("DbPort").MustString("3306")
	DbUser = file.Section("dataBase").Key("DbUser").MustString("root")
	DbPassword = file.Section("dataBase").Key("DbPassword").MustString("yaoyuan52163#")
	DbName = file.Section("dataBase").Key("DbName").MustString("ginblog")
}


func LoadQiNiu(file *ini.File){
	AccessKey = file.Section("qiNiu").Key("AccessKey").String()
	AccessKey = file.Section("qiNiu").Key("AccessKey").String()
	Bucket = file.Section("qiNiu").Key("Bucket").String()
	QiniuServer = file.Section("qiNiu").Key("QiniuServer").String()
}