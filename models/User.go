package models

import (
	db "employee_manage/config"
	modelErrors "employee_manage/constant"

	"time"
)

type User struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	EmployeeCode string    `json:"employee_code"`
	FullName     string    `json:"full_name"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Gender       bool      `json:"gender"`
	Birthday     time.Time `json:"birthday"`
	Address      string    `json:"address"`
	Password     string    `json:"password"`
	DepartmentID int       `gorm:"not null;column:department_id" json:"department_id"`
	RoleID       int       `gorm:"not null;column:role_id" json:"role_id"`
	Department   Department
	Role         Role
	CreatedAt    time.Time `json:"created_at,omitempty" gorm:"autoCreateTime:mili"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime:mili"`
}

func CreateUser(user *User) error {
	newUser := db.DB.Create(&user)
	return newUser.Error
}

func (u *User) GetAllUser(users *[]User, limit int, offset int) (err error) {
	err = db.DB.Limit(limit).Offset(offset).Find(users).Error

	if err != nil {
		return err
	}

	return
}

func GetUserByID(user *User, id int) error {
	err := db.DB.First(user, id).Error
	return err
}

func UpdateUserByID(id int, userMap map[string]interface{}) (user User, err error) {
	user.ID = id
	err = db.DB.Model(&user).
		Select("full_name", "phone", "email", "gender", "birthday", "address", "department_id").
		Updates(userMap).
		Error

	if err != nil {
		return User{}, err
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
