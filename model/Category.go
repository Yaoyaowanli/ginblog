package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
	"log"
)

//Category 文章类别
type Category struct {
	Id   uint   `gorm:"primary_key;auto_increment" json:"id"`   //类别id（主键约束 特征：not null + unique(不可重复)）
	                               // auto_increment (自动维护生成唯一主键)
	Name string `gorm:"type:varchar(32);not null" json:"name"`	//类别名
}

//CheckCategory 查询分类是否存在
func CheckCategory (id uint) int {
	var cate Category
	db.Select("id").Where("id=?",id).First(&cate)
	if cate.Id>0{
		return errmsg.SUCCESS
	}
	return errmsg.ERROR_CATE_NOT_EXIST
}

//CreateCate 新增分类
func CreateCate(data *Category)int{
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//GetCate 查询分类列表
func GetCate(pageSize,pageNo int)([]Category,int64){
	var cate []Category
	var total int64
	offset := (pageNo-1)*pageSize
	if pageNo == -1 && pageSize == -1{
		offset = -1
	}
	err = db.Limit(pageSize).Offset(offset).Find(&cate).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,0
	}
	return cate,total
}

//EditCate 编辑分类
func EditCate(id int,data *Category)int{
	maps := make(map[string]interface{})
	maps["id"]=data.Id
	maps["name"]=data.Name
	err = db.Model(data).Where("id=?",id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func FindByCateId(id int) *Category{
	var cate *Category
	err = db.Last(&cate,id).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return cate
}

//DeleteCate 删除类别
func DeleteCate(id int)int{
	var cate *Category
	err = db.Where("id=?",id).Delete(&cate).Error
	if err != nil {
		log.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
