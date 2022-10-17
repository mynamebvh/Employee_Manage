package services

import (
	modelErrors "employee_manage/constant"
	"employee_manage/models/dto"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ExportExcel(deName string, users []dto.QueryGetUsersByDepartmentID) (err error) {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "STT")
	f.SetCellValue("Sheet1", "B1", "Họ tên")
	f.SetCellValue("Sheet1", "C1", "Mã nhân viên")
	f.SetCellValue("Sheet1", "D1", "Số điện thoại")
	f.SetCellValue("Sheet1", "E1", "Email")
	f.SetCellValue("Sheet1", "F1", "Giới tính")
	f.SetCellValue("Sheet1", "G1", "Địa chỉ")
	f.SetCellValue("Sheet1", "H1", "Phòng")

	index := 2
	for _, user := range users {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", index), index-1)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", index), user.FullName)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", index), user.EmployeeCode)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", index), user.Phone)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", index), user.Email)
		if user.Gender {
			f.SetCellValue("Sheet1", fmt.Sprintf("F%d", index), "Nam")
		} else {
			f.SetCellValue("Sheet1", fmt.Sprintf("F%d", index), "Nữ")
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", index), user.Address)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", index), user.Name)
		index++
	}

	if err := f.SaveAs(fmt.Sprintf("./files/%s.xlsx", deName)); err != nil {
		return modelErrors.NewAppError(err, modelErrors.ValidationError)
	}

	return
}
