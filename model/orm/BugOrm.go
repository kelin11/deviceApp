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


// 获取所有的故障表数据--无条件查询
func GetAllBugTable() ([]model.Bug, int){
	var bugList []model.Bug
	err := db.Find(&bugList)
	if err != nil{
		log.Fatal("Function GetAllBugTable sql查询错误: ", err)
		return nil, errmsg.ERROR
	}
	return bugList, errmsg.SUCCESS
}

// 根据实验室Lab获取故障表数据--条件查询 key=Lab
func GetBugTableByLab(lab string) (int, []model.Bug){
	var bugList []model.Bug
	err := db.Where("lab = ?", lab).Find(&bugList).Error
	if err != nil {
		return errmsg.ERROR, nil
	}
	return errmsg.SUCCESS, bugList
}