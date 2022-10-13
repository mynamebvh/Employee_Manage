package models

import (
	db "employee_manage/config"
	modelErrors "employee_manage/constant"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

var (
	TEST_IMPORT = "TEST"
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
