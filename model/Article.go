package model

import "gorm.io/gorm"

//Article 文章
type Article struct {
							//foreignKey（外键约束：只能出现Category中主键id允许的值，约束文章的cid字段只能出现文章类别中id允许的值）
	Category Category `gorm:"foreignKey:Cid"`     //Category 文章类别
	gorm.Model
	Title string  `gorm:"type:varchar(100);not null" json:"title"`  //Title 文章标题
	Cid int `gorm:"type:int;not null" json:"cid"`			//分类id
	Desc string `gorm:"type:varchar(200)" json:"desc"`		//描述
	Content string `gorm:"type:longtext" json:"content"`	//Content 内容
	Img string `gorm:"type:varchar(100)" json:"img"`		//Img 图片
}
