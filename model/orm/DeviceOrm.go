package orm

import (
	errmsg "deviceApp/code"
	"deviceApp/model"
	"log"
)

//查看所有不可用设备 返回的是 设备的 id name 以及所属于的实验室id
func GetAllUnDevice() ([]model.Device, int) {
	var udList []model.Device
	// err := db.Where("status = 0").Select("d_id", "d_name", "lab", "bug_num").Find(&udList).Error
	err := db.Where("status = 0").Find(&udList).Error
	if err != nil {
		log.Fatal("查询不可用仪器出现问题:", err)
		return nil, errmsg.ERROR
	}

	return udList, errmsg.SUCCESS

}

// 获取某实验室中所有的设备
func GetAllDevice(lab string) ([]model.Device, int) {
	var deviceList []model.Device
	err := db.Where("lab = ?", lab).Select("d_id", "d_name", "lab").Find(&deviceList).Error
	if err != nil {
		log.Fatal("Function GetAllDevice sql查询错误: ", err)
		return nil, errmsg.ERROR
	}
	return deviceList, errmsg.SUCCESS
}
