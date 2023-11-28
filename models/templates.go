package models

import "github.com/jinzhu/gorm"

type Templates struct {
	gorm.Model
	Date   string `gorm:"column:date;type: Date"`   // 2016-02-12这种格式的日期
	Detail string `gorm:"column:detail;type: JSON"` // 字符串数组
	Day    int8   `gorm:"column:day;type: tinyInt"` // 1～7，分别对应周一到周日
	UserID int    `gorm:"column:userid"`

	User Users `gorm:"foreignkey:UserID;references:ID"`
}
