package model

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	D_id   string `gorm:"varchar(20);not null" json:"device_id"`
	D_name string `gorm:"varchar(20);not null" json:"device_name"`
	Lab    string `gorm:"varchar(20);not null" json:"lab"`
	Status int    `gorm:"int;not null" json:"status"`
}
