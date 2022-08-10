package model

import "gorm.io/gorm"

//故障上报
type Bug struct {
	gorm.Model
	Lab    string `gorm:"type:varchar(20);not null" json:"lab"`
	D_name string `gorm:"type:varchar(20);not null" json:"device_name"`
	D_id   string `gorm:"type:varchar(20);not null" json:"device_id"`

	U_name   string `gorm:"type:varchar(20);not null" json:"username"`
	Img      string `gorm:"type:varchar(100)" json:"image"`
	Comments string `gorm:"type:text" json:"comments"`
}
