package model

import "gorm.io/gorm"

type BugDetail struct {
	gorm.Model
	Lab      string `gorm:"type:varchar(20);not null" json:"lab" xlsx:"lab"`
	D_name   string `gorm:"type:varchar(20);not null" json:"device_name" xlsx:"device_name"`
	D_id     string `gorm:"type:varchar(20);not null" json:"device_id" xlsx:"device_id"`
	Bug_num  int    `gorm:"varchar(20);not null" json:"bug_num" xlsx:"bug_num"`
	Comments string `gorm:"type:text" json:"comments" xlsx:"comments"`
}
