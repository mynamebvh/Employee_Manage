package models

import (
	db "employee_manage/config"
	modelErrors "employee_manage/constant"
	"encoding/json"
	"time"
)

type Token struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Value     string    `json:"value"`
	Type      string    `json:"type"`
	UserID    int       `gorm:"not null;column:user_id" json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime:mili"`
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
