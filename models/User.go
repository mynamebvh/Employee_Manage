package models

import (
	db "employee_manage/config"
	modelErrors "employee_manage/constant"
	"employee_manage/models/dto"
	"employee_manage/utils"
	"encoding/json"
	"errors"

	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	EmployeeCode string    `json:"employee_code" gorm:"not null"`
	FullName     string    `json:"full_name" gorm:"not null;colum:full_name"`
	Phone        string    `json:"phone" gorm:"not null;unique"`
	Email        string    `json:"email" gorm:"not null;unique"`
	Gender       bool      `json:"gender" gorm:"not null"`
	Birthday     time.Time `json:"birthday"`
	Address      string    `json:"address" gorm:"not null"`
	Password     string    `json:"-" gorm:"not null"`
	DepartmentID int       `json:"department_id" gorm:"not null;column:department_id" `
	RoleID       int       `json:"-" gorm:"not null;column:role_id" `
	Token        Token     `json:"-"`
	CreatedAt    time.Time `json:"-" gorm:"autoCreateTime:mili"`
	UpdatedAt    time.Time `json:"-" gorm:"autoUpdateTime:mili"`
}

/* HOOK */
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := HashPassword(u.Password)
	u.Password = hash

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {

	if tx.Statement.Changed("Password") {
		newPassword := tx.Statement.Dest.(map[string]interface{})["password"].(string)
		hash, _ := HashPassword(newPassword)
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

func GetUsers(c *gin.Context) (data dto.QueryPagination, err error) {
	var user []dto.QueryUserJoin
	var page, limit int
	err = db.DB.Model(&User{}).Scopes(utils.Paginate(c, &page, &limit)).
		Select(`
		users.id, 
		users.full_name, 
		users.employee_code, 
		users.phone, 
		users.email,
		users.gender,
		users.birthday,
		users.address,
		d.name department_name`).
		Joins("left join departments d on users.department_id = d.id").
		Scan(&user).Error

	if err != nil {
		return
	}

	var count int64
	db.DB.Model(&User{}).Count(&count)

	data = dto.QueryPagination{
		Current:  page,
		Total:    count,
		PageSize: limit,
		Data:     user,
	}

	return
}

func GetMe(user *User, id int) (queryResult dto.QueryUserJoin, err error) {
	db.DB.Model(user).
		Select(
			"users.id",
			"users.full_name",
			"users.employee_code",
			"users.phone",
			"users.email",
			"users.gender",
			"users.address",
			"departments.name as department_name",
		).
		Where("users.id = ?", id).
		Joins("left join departments on users.department_id = departments.id").Scan(&queryResult)

	return
}

func GetRole(user *User, id int) (role dto.QueryCheckManageAccess, err error) {
	db.DB.Model(user).
		Select(
			"ud.department_id",
			"name",
		).
		Where("users.id = ?", id).
		Joins("left join user_departments as ud on users.id = ud.user_id").
		Joins("left join roles on users.role_id = roles.id").
		Scan(&role)

	return
}

func CheckUserInDepartment(managerID int, userID int) bool {
	var id int
	queryGetDepartmentID := db.DB.Model(UserDepartment{}).
		Select("department_id").
		Where("user_id = ?", managerID)

	db.DB.Model(User{}).
		Select("id").
		Where("id = ? AND department_id = (?)", userID, queryGetDepartmentID).
		Scan(&id)

	return id != 0
}

func GetUserByID(id int) (user User, err error) {
	err = db.DB.First(&user, id).Error

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

func GetUserByEmail(email string) (user User, err error) {
	err = db.DB.Where("email=?", email).First(&user).Error

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

	if err != nil {
		err = modelErrors.NewAppErrorWithType(modelErrors.NotFound)
		return
	}

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

func ResetPassword(email string, newPassword string) (err error) {
	err = db.DB.Model(&User{}).Where("email = ?", email).Update("password", newPassword).Error
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
