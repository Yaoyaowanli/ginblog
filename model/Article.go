package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
	"log"
)

//Article 文章
type Article struct {
							//foreignKey（外键约束：只能出现Category中主键id允许的值，约束文章的cid字段只能出现文章类别中id允许的值）
	Category Category     `gorm:"foreignkey:Cid"`     //Category 文章类别
	gorm.Model
	Title string  `gorm:"type:varchar(100);not null" json:"title"`  //Title 文章标题
	Cid int `gorm:"type:int;not null;" json:"Cid"`			//分类id
	Desc string `gorm:"type:varchar(200)" json:"desc"`		//描述
	Content string `gorm:"type:longtext" json:"content"`	//Content 内容
	Img string `gorm:"type:varchar(100)" json:"img"`		//Img 图片
}

//CreateArt 新增文章
func CreateArt(data *Article) int {
	err = db.Create(&data).Error
	if err != nil {
		log.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询分类下所有文章


// todo 查询单个文章


// todo  查询文章列表
func GetArt(pageSize,pageNo int)([]Article,int){
	var art []Article
	offset := (pageNo-1)*pageSize
	if pageNo == -1 && pageSize == -1{
		offset = -1
	}
	err := db.Preload("Category").Limit(pageSize).Offset(offset).Find(&art).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,errmsg.ERROR
	}
	return art,errmsg.SUCCESS
}

// EditArt  编辑文章
func EditArt(id int,data *Article)int{
	var art Article
	maps := make(map[string]interface{})
	maps["Category"]=data.Category
	maps["title"]=data.Title
	maps["Cid"]=data.Cid
	maps["desc"]=data.Desc
	maps["content"]=data.Content
	maps["img"]=data.Img
	err = db.Model(&art).Where("id=?",id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//DeleteArt 删除文章
func DeleteArt(id int)int {
	var art Article
	err = db.Where("id=?",id).Delete(art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}