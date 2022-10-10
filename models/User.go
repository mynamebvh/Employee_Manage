package models

import (
	db "employee_manage/config"
	modelErrors "employee_manage/constant"
	"encoding/json"

	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	EmployeeCode string    `json:"employee_code" gorm:"not null"`
	FullName     string    `json:"full_name" gorm:"not null"`
	Phone        string    `json:"phone" gorm:"not null;unique"`
	Email        string    `json:"email" gorm:"not null;unique"`
	Gender       bool      `json:"gender" gorm:"not null"`
	Birthday     time.Time `json:"birthday" gorm:"not null"`
	Address      string    `json:"address" gorm:"not null"`
	Password     string    `json:"password" gorm:"not null"`
	DepartmentID int       `json:"department_id" gorm:"not null;column:department_id" `
	RoleID       int       `json:"role_id" gorm:"not null;column:role_id" `
	// Department   Department
	// Role         Role
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime:mili"`
}

func CreateUser(user *User) (err error) {
	// newUser := db.DB.Create(&user)
	// return newUser.Error

	err = db.DB.Create(user).Error

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

func (u *User) GetAllUser(users *[]User, limit int, offset int) (err error) {
	err = db.DB.Limit(limit).Offset(offset).Find(users).Error

	if err != nil {
		return err
	}

	return
}

func GetUserByID(user *User, id int) (err error) {
	err = db.DB.First(user, id).Error

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

func UpdateUserByID(id int, userMap map[string]interface{}) (user User, err error) {
	user.ID = id
	err = db.DB.Model(&user).
		Select("full_name", "phone", "email", "gender", "birthday", "address", "department_id").
		Updates(userMap).
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

	err = db.DB.Where("id = ?", id).First(&user).Error
	return
}

func DeleteUserByID(id int) (err error) {
	tx := db.DB.Delete(&User{}, id)

	if tx.Error != nil {
		err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = modelErrors.NewAppErrorWithType(modelErrors.NotFound)
	}

	return
}
