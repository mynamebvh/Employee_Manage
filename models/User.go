package models

import (
	db "employee_manage/config"
	modelErrors "employee_manage/constant"
	"employee_manage/utils"
	"encoding/json"
	"errors"

	"time"

	"golang.org/x/crypto/bcrypt"
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
	Password     string    `json:"-" gorm:"not null"`
	DepartmentID int       `json:"department_id" gorm:"not null;column:department_id" `
	RoleID       int       `json:"role_id" gorm:"not null;column:role_id" `
	Token        Token     `json:"-"`
	CreatedAt    time.Time `json:"created_at,omitempty" gorm:"autoCreateTime:mili"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime:mili"`
}

/* HOOK */
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := HashPassword(u.Password)
	u.Password = hash

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	hash, err := HashPassword(u.Password)

	if tx.Statement.Changed("Password") {
		tx.Statement.SetColumn("password", hash)
	}

	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func CreateUser(user *User) (err error) {
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

func GetUserByEmail(user *User, email string) (err error) {
	err = db.DB.Where("email=?", email).First(user).Error

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

func ChangePassword(user User, oldPassword string, newPassword string) (err error) {

	isCorrect := CheckPasswordHash(user.Password, oldPassword)

	if !isCorrect {
		err = modelErrors.NewAppError(errors.New("old password not correct"), "Validation")
		return
	}

	err = db.DB.Model(&user).Update("password", newPassword).Error

	return
}

func ResetPassword(token Token, newPassword string) (err error) {
	var user User
	err = db.DB.Model(&user).Where("email = ?", token.Email).Update("password", newPassword).Error
	return
}

func GenerateResetPassword(user User, email string) (code string, err error) {

	code = utils.RandomStringSecret(20)

	expired := time.Now().Local().Add(time.Minute * time.Duration(5))
	err = db.DB.Create(&Token{
		Value:   code,
		Type:    "reset_password",
		Email:   email,
		Expired: &expired,
	}).Error

	return
}
