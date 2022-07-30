package orm

import (
	errmsg "deviceApp/code"
	"deviceApp/model"
	"log"
)

// 获取所有的维修反馈表数据
func GetAllSolutionTable() (int, []model.Solution) {
	var solutionList []model.Solution
	err = db.Find(&solutionList).Error
	if err != nil{
		log.Fatal("Function GetAllSolutionTable sql查询失败: ", err)
		return errmsg.ERROR, nil
	}
	return errmsg.SUCCESS, solutionList
}

// 维修人员上传反馈
func SubmitSolutionData(solution model.Solution) int {
	err := db.Create(&solution).Error
	if err != nil{
		log.Fatal("Function SubmitSolutionData sql插入失败: ", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}