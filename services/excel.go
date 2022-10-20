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

func ExportExcelWorkingTime(month int, data []dto.QueryGetWorkingTimeInMonth) (err error) {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "STT")
	f.SetCellValue("Sheet1", "B1", "Họ tên")
	f.SetCellValue("Sheet1", "C1", "Mã nhân viên")
	f.SetCellValue("Sheet1", "D1", "Phòng")
	f.SetCellValue("Sheet1", "E1", "Số ngày làm việc")
	f.SetCellValue("Sheet1", "F1", "Số giờ làm việc")

	index := 2
	for _, u := range data {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", index), index-1)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", index), u.FullName)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", index), u.EmployeeCode)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", index), u.DepartmentName)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", index), u.Days)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", index), u.Hours)
		index++
	}

	if err := f.SaveAs(fmt.Sprintf("./files/working-time%d.xlsx", month)); err != nil {
		return modelErrors.NewAppError(err, modelErrors.ValidationError)
	}

	return
}

type DataQueue struct {
	ID       string
	FullName string
	Email    string
}

func ReadExcelResetPassword(path string) (data []DataQueue, err error) {

	f, err := excelize.OpenFile(path)
	if err != nil {
		err = modelErrors.NewAppError(err, modelErrors.ValidationError)
		return
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		err = modelErrors.NewAppError(err, modelErrors.ValidationError)
		return
	}

	for i := 1; i < len(rows); i++ {
		data = append(data, DataQueue{
			ID:       rows[i][0],
			FullName: rows[i][1],
			Email:    rows[i][2],
		})
		fmt.Println()
	}

	return
}
