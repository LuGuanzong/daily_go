package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Templates struct {
	gorm.Model
	Date   time.Time `json:"date" gorm:"column:date;"`     // 2016-02-12这种格式的日期
	Detail string    `json:"detail" gorm:"column:detail;"` // 字符串数组
	Day    int8      `json:"day" gorm:"column:day;"`       // 1～7，分别对应周一到周日
	UserID int       `json:"userID" gorm:"column:userid"`

	User Users `gorm:"foreignkey:UserID;references:ID"`
}
