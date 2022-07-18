package model

import (
	"time"

	"gorm.io/gorm"
)

type Solution struct {
	gorm.Model
	D_id     string
	D_name   string
	Lab    	 string
	B_id     int
	S_time   time.Time
	Feedback string
}
