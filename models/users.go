package models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Phone    string `gorm:"column:phone;size:20;unique:true"`
	Email    string `gorm:"column:email;size:256;unique:true"`
	Password string `gorm:"column:password"`
}
