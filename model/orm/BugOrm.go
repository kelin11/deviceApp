package orm

import (
	errmsg "deviceApp/code"
	"deviceApp/model"
	"fmt"
	"log"

	"gorm.io/gorm"
)

//data 故障数据 uName 学号  这里是传递用户id还是用户名称?
func SubmitUnDevice(data model.Bug, uName string) (code int) {
	data.U_name = uName
	err := db.Model(&model.Device{}).
		Where("d_id = ?", data.D_id).
		Update("status", 0).
		Update("bug_num", gorm.Expr("bug_num + ?", 1)).Error
	err = db.Create(&data).Error
	if err != nil {
		log.Fatal("插入数据失败")
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 获取所有的故障表数据
//return :
func GetBugList() ([]model.Bug, int) {
	var bugList []model.Bug
	err := db.Select("d_name", "d_id", "lab").Find(&bugList)
	if err != nil {
		log.Fatal("Function GetAllBugTable sql查询错误: ", err)
		return nil, errmsg.ERROR
	}
	return bugList, errmsg.SUCCESS
}

//查询bug详细信息 以此为基准输出excel文件
func GetBugDetail(bugId string) (model.BugDetail, int) {
	var bugDetail model.BugDetail
	var bug model.Bug
	//var device model.Device
	db.Where("id = ?", bugId).Select("d_id").First(&bug)
	fmt.Println(bug.D_id)
	err = db.Model(&bug).
		Select("bugs.*,devices.bug_num").
		Joins("left join devices on devices.d_id=bugs.d_id").
		Where("bugs.d_id = ?", bug.D_id).
		Scan(&bugDetail).Error
	if err != nil {
		log.Fatal("bugDetail查询出现问题", err)
		return bugDetail, errmsg.ERROR
	}
	return bugDetail, errmsg.SUCCESS

}

func ExportBugExcel() {

}