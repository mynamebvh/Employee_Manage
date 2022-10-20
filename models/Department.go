package models

import (
	db "employee_manage/config"
	modelErrors "employee_manage/constant"
	"employee_manage/models/dto"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"not null;unique"`
	DepartmentCode string    `json:"department_code" gorm:"not null;unique"`
	Address        string    `json:"address"`
	Status         bool      `json:"status" gorm:"not null"`
	User           User      `json:"-"`
	CreatedAt      time.Time `json:"-" gorm:"autoCreateTime:mili"`
	UpdatedAt      time.Time `json:"-" gorm:"autoUpdateTime:mili"`
}

func GetUsersByDepartmentID(id int) (users []dto.QueryGetUsersByDepartmentID, err error) {
	err = db.DB.Table("users as u").
		Select("u.full_name", "u.employee_code", "u.phone", "u.email", "u.gender", "u.address", "d.name").
		Joins(`inner join departments as d on u.department_id = d.id 
			AND u.department_id = ?
			`, id).
		Where("u.role_id = 3").
		Scan(&users).Error

	if err != nil {
		return
	}

	return
}

func GetDepartmentByRequestID(id int) (user User, err error) {
	var request Request

	err = db.DB.Model(&request).
		Select("users.*").
		Joins("left join users on requests.user_id = users.id and requests.id = ?", id).
		Scan(&user).Error
	return
}

func CreateDepartment(de *Department) (err error) {
	err = db.DB.Create(de).Error

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError modelErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return err
		}
		switch newError.Number {
		case 1062:
			err = modelErrors.NewAppErrorWithType(modelErrors.ResourceAlreadyExists)
			return

		default:
			err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		}
	}

	return
}

func GetDepartmentByID(de *Department, id int) (err error) {
	err = db.DB.First(de, id).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = modelErrors.NewAppErrorWithType(modelErrors.NotFound)
		default:
			err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		}
	}

	return
}

func UpdateDepartmentByID(id int, departmentMap map[string]interface{}) (department Department, err error) {
	department.ID = id
	err = db.DB.Model(&department).
		Updates(departmentMap).
		Error

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError modelErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return
		}
		switch newError.Number {
		case 1062:
			err = modelErrors.NewAppErrorWithType(modelErrors.ResourceAlreadyExists)
			return

		default:
			err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		}
	}

	err = db.DB.Where("id = ?", id).First(&department).Error
	return
}

func DeleteDepartmentByID(id int) (err error) {
	tx := db.DB.Delete(&Department{}, id)

	if tx.Error != nil {
		err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = modelErrors.NewAppErrorWithType(modelErrors.NotFound)
	}

	return
}

func GetWorkingTimeInMonth(month int) (q []dto.QueryGetWorkingTimeInMonth, err error) {

	err = db.DB.Raw(`SELECT
	r.user_id,
	users.full_name,
	users.employee_code,
	SEC_TO_TIME(r.sec) as hours,
	r.days as days,
	d.name as department_name
	FROM
	(
	SELECT
		user_id,
		SUM(
			TIME_TO_SEC(
			TIMEDIFF( checkout_time, checkin_time ))) AS sec,
		COUNT(user_id) as days
	FROM
		calendars AS c
	WHERE MONTH ( checkin_time ) = ?
	GROUP BY
	user_id 
	) AS r
	LEFT JOIN users on r.user_id = users.id
	LEFT JOIN departments d on 	users.department_id = d.id`, month).Scan(&q).Error

	if err != nil {
		err = modelErrors.NewAppError(err, modelErrors.ValidationError)
		return
	}
	return
}
