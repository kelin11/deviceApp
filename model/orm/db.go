package orm

import (
	"database/sql"
	"deviceApp/model"
	"deviceApp/settings"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	//root:root@tcp(localhost:3306)/deviceApp?charset=utf8&parseTime=true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		settings.DbUser,
		settings.DbPassWord,
		settings.DbHost,
		settings.DbPort,
		settings.DbName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var sqlDB *sql.DB
	sqlDB, err = db.DB()

	//自动保持数据库处于最新状态
	db.AutoMigrate(&model.User{}, &model.Device{}, &model.Solution{}, &model.Bug{})

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
