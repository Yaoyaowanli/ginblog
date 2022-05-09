package model

import "gorm.io/gorm"

//Category 文章类别
type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(32);not null" json:"name"`	//类别名
	Id   uint   `gorm:"primary_key;auto_increment" json:"id"`   //类别id（主键约束 特征：not null + unique(不可重复)）
								//auto_increment (自动维护生成唯一主键)
}
