package orm

import (
	errmsg "deviceApp/code"
	"deviceApp/model"
	"log"
)

// 获取所有的维修反馈表数据列表
func GetAllSolutionTable() (int, []model.Solution) {
	var solutionList []model.Solution
	// err = db.Select("d_id", "d_name", "lab").Find(&solutionList).Error
	err = db.Find(&solutionList).Error
	if err != nil {
		log.Fatal("Function GetAllSolutionTable sql查询失败: ", err)
		return errmsg.ERROR, nil
	}
	return errmsg.SUCCESS, solutionList
}

// 维修人员上传反馈(存疑问) 是excel转到数据库?
func SubmitSolutionData(solution model.Solution) int {
	//改变设备状态statuscc 为可用
	err = db.Model(&model.Device{}).Where("d_id = ?", solution.D_id).Update("status", 0).Error

	err = db.Create(&solution).Error

	if err != nil {
		log.Fatal("Function SubmitSolutionData sql插入失败: ", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
