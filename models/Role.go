package models

import (
	modelErrors "employee_manage/constant"
)

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	User User
}

func CheckManageAccess(id int, departmentID int, model string) (isNext bool, err error) {
	var user User

	if model == "requests" {
		if user, err = GetDepartmentByRequestID(id); err != nil {
			err = modelErrors.NewAppError(err, modelErrors.ValidationError)
			return
		}

	} else if model == "users" {

		if user, err = GetUserByID(id); err != nil {
			err = modelErrors.NewAppError(err, modelErrors.ValidationError)
			return
		}
	} else {
		isNext = true
		return
	}

	if user.DepartmentID == departmentID {
		isNext = true
		return
	}

	isNext = false
	return
}
