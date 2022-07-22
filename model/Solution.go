package model

import (
	"gorm.io/gorm"
)

//维修方案
type Solution struct {
	gorm.Model
	D_id     string `gorm:"type:varchar(20);not null" json:"device_id"`
	D_name   string `gorm:"type:varchar(20);not null" json:"device_name"`
	U_name   string `gorm:"type:varchar(20);not null" json:"username"`
	Lab      string `gorm:"type:varchar(20);not null" json:"lab"`
	B_id     int    `gorm:"type:int" json:"bug_id"`
	Feedback string `gorm:"type:text" json:"feedback"`
}
