package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name;unique;index;not null;comment:'用户昵称'"`
	Phone    string `json:"phone" gorm:"column:phone;size:20;unique;comment:'手机号'"`
	Email    string `json:"email" gorm:"column:email;size:256;unique;comment:'邮箱'"`
	Password string `json:"password" gorm:"column:password;comment:'密码'"`
}
