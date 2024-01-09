package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name;unique;index"`
	Phone    string `json:"phone" gorm:"column:phone;size:20;unique"`
	Email    string `json:"email" gorm:"column:email;size:256;unique"`
	Password string `json:"password" gorm:"column:password"`
}
