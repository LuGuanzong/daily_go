package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Date   time.Time       `json:"date" gorm:"column:date;not null;comment:'日期，仅精确到天'"`
	Detail json.RawMessage `json:"detail" gorm:"column:detail;not null;comment:'任务'"`
	Userid int             `json:"userid"  gorm:"column:userid;not null;comment:'用户id'"`

	User User `gorm:"foreignkey:Userid;references:ID"`
}
