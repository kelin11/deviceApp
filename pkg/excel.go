package pkg

import (
	"bytes"
	"deviceApp/model"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
)

func DownExcel(bugDetails []model.BugDetail) (*bytes.Buffer, error) {
	newSheetName := "default"
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", newSheetName)
	rowNum := 1

	// save
	for  i, bugDetail := range bugDetails {
		// reflect
		p := reflect.TypeOf(bugDetail)
		// title
		if i == 0 {
			fmt.Println(p.String())
			// 表头
			header := make([]string, 0)
			for j := 0; j < p.NumField(); j++ {
				//  fieldName of Struct
				key := p.Field(j)
				fmt.Println("name=", key.Name, "tag=", key.Tag.Get("xlsx"))
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
		for j := 0; j < p.NumField(); j++ {
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
	if err := f.SaveAs("./cmd/deviceapp-demo/Book2.xlsx"); err != nil {
		fmt.Println(err)
	}
	return f.WriteToBuffer()
}