package model

import "gorm.io/gorm"

type BugDetail struct {
	gorm.Model
	Lab      string `gorm:"type:varchar(20);not null" json:"lab"`
	D_name   string `gorm:"type:varchar(20);not null" json:"device_name"`
	D_id     string `gorm:"type:varchar(20);not null" json:"device_id"`
	Bug_num  int    `gorm:"varchar(20);not null" json:"bug_num"`
	Comments string `gorm:"type:text" json:"comments"`
}
