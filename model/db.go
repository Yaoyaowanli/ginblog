package model

//使用gorm操作数据库

import (
	"fmt"
	"ginblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var err error



func InitDb(){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
					utils.DbUser,utils.DbPassword,utils.DbHost,utils.DbPort,utils.DbName)
	db,err = gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败，请检查参数是否正确,err:",err)
	}


	db.AutoMigrate(&User{},&Article{},&Category{})

	sqlDb,_:=db.DB()
	//设置最大可连接数
	sqlDb.SetMaxOpenConns(100)
	//连接池最大允许的空闲连接数
	sqlDb.SetMaxIdleConns(10)
	//设置连接最大的可复用时间
	sqlDb.SetConnMaxLifetime(10*time.Second)

}
