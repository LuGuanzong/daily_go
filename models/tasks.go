package models

import "github.com/jinzhu/gorm"

// detail结构如下
// 这是json数组
//{
//label: '喝奶茶'；
//done: 1
//}

type Tasks struct {
	gorm.Model
	Date   string `gorm:"column:date;type: Date"`
	Detail string `gorm:"column:detail;type: JSON"`
	UserID int    `gorm:"column:userid"`

	User Users `gorm:"foreignkey:UserID;references:ID"`
}
