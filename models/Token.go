package models

import (
	db "employee_manage/config"
	modelErrors "employee_manage/constant"
	"encoding/json"
	"errors"
	"time"
)

type Token struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Value     string     `json:"value"`
	Type      string     `json:"type"`
	Email     string     `json:"email"`
	Expired   *time.Time `json:"expired"`
	UserID    *int       `gorm:"column:user_id" json:"user_id"`
	CreatedAt time.Time  `json:"created_at,omitempty" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time  `json:"-" gorm:"autoUpdateTime:mili"`
}

func GetTokenByValueAndType(value string, typeToken string) (token Token, err error) {
	err = db.DB.Where("value = ? AND type = ?", value, typeToken).Find(&token).Error

	if token.ID == 0 {
		err = modelErrors.NewAppError(errors.New("refresh token not found"), modelErrors.NotFound)
		return
	}

	return
}

func CreateToken(token *Token) (err error) {
	err = db.DB.Create(token).Error

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

func CheckToken(code string) (token Token, err error) {
	err = db.DB.Where("value = ?", code).Find(&token).Error

	if token.ID == 0 {
		err = modelErrors.NewAppError(errors.New("code not valid"), "NOT_VALID")
		return
	}

	if time.Now().After(*token.Expired) {
		err = modelErrors.NewAppError(errors.New("code expired"), "NOT_VALID")
		return
	}

	return
}

func DeleteToken(tokenString string) (err error) {
	var token Token
	tx := db.DB.Where("value = ?", tokenString).Delete(&token)

	if tx.Error != nil {
		err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = modelErrors.NewAppErrorWithType(modelErrors.NotFound)
	}

	return
}
