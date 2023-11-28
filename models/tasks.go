package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// detail结构如下
// 这是json数组
//{
//label: '喝奶茶'；
//done: 1
//}

type Tasks struct {
	gorm.Model
	Date   time.Time `json:"date" gorm:"column:date;"`
	Detail string    `json:"detail" gorm:"column:detail;"`
	UserID int       `json:"userID"  gorm:"column:userid;"`

	User Users `gorm:"foreignkey:UserID;references:ID"`
}
