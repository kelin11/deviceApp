package model

import "gorm.io/gorm"

//记录设备
type Device struct {
	gorm.Model
	D_id   string `gorm:"varchar(20);not null" json:"device_id"`
	D_name string `gorm:"varchar(20);not null" json:"device_name"`
	Lab    string `gorm:"varchar(20);not null" json:"lab"`
	//status : 0不可用 1 可用 2 维修中
	Status int `gorm:"int;not null" json:"status"`
}
