package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Date   time.Time `json:"date" gorm:"column:date;"`
	Detail string    `json:"detail" gorm:"column:detail;"`
	Userid int       `json:"userid"  gorm:"column:userid;"`

	User User `gorm:"foreignkey:Userid;references:ID"`
}
