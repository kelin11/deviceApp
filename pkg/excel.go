package pkg

import (
	"bytes"
	"deviceApp/model"
	"fmt"
	"reflect"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func DownExcel(bugDetails []model.BugDetail) (*bytes.Buffer, error) {
	newSheetName := "故障表"
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", newSheetName)
	rowNum := 1

	// save
	for i, bugDetail := range bugDetails {
		fmt.Println("===============")
		fmt.Println(bugDetail)
		// reflect
		p := reflect.TypeOf(bugDetail)
		// title
		if i == 0 {
			// 表头
			header := make([]string, 0)
			fmt.Println("================Numfield", p.NumField())
			for j := 1; j < p.NumField(); j++ {
				//  fieldName of Struct
				key := p.Field(j)
				fmt.Println("name =", key.Name, "tag =", key.Tag.Get("xlsx"))
				tag := key.Tag.Get("xlsx")
				if tag != "" {
					header = append(header, tag)
				}
			}
			f.SetSheetRow(newSheetName, "A1", &header)
		}

		valObj := reflect.ValueOf(bugDetail)
		// read data to excel
		valueList := make([]interface{}, 0)
		for j := 1; j < p.NumField(); j++ {
			key := p.Field(j)
			tag := key.Tag.Get("xlsx")
			if tag != "" {
				value := valObj.Field(j).Interface()
				valueList = append(valueList, value)
			}
		}
		rowNum++
		f.SetSheetRow(newSheetName, "A"+strconv.Itoa(rowNum), &valueList)
	}
	// 本地测试保存代码
	// if err := f.SaveAs("Book2.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }
	return f.WriteToBuffer()
}
