package orm

import (
	errmsg "deviceApp/code"
	"deviceApp/model"
	"log"
)

//data 故障数据 uName 学号  这里是传递用户id还是用户名称?
func SubmitUnDevice(data model.Bug, uName string) (code int) {
	data.U_name = uName
	err := db.Create(&data).Error
	if err != nil {
		log.Fatal("插入数据失败")
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
