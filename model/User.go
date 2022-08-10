package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null" json:"username" `
	PassWord string `gorm:"type:varchar(20);not null" json:"password" `
	//0 维修人员 1 普通学生 2 后台管理员
	Role int `gorm:"type:int;not null" json:"role"`
}

