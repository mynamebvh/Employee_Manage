package models

import (
	db "employee_manage/config"
	modelErrors "employee_manage/constant"
	"encoding/json"
	"errors"
	"time"
)

type Calendar struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	UserID       int       `gorm:"not null;column:user_id" json:"user_id"`
	User         User      `json:"-"`
	CheckinTime  time.Time `json:"checkin_time" gorm:"autoCreateTime:mili"`
	CheckoutTime time.Time `json:"checkout_time" gorm:"default:(-)"`
}

func CheckValidCheckin(userID int) (err error) {
	var count int
	db.DB.
		Raw(
			`SELECT COUNT(id) 
			FROM calendars 
			WHERE calendars.user_id = ?
			AND DATEDIFF(calendars.checkin_time, "2022-10-17 14:37:43.560") = 0`, userID).Scan(&count)

	if count > 0 {
		err = errors.New("have you checkin in today?")
		return
	}

	return
}

func CheckValidCheckout(userID int) (id int, err error) {
	timeNow := time.Now().Format("2006-01-02 15:04:05.000")

	db.DB.
		Raw(
			`SELECT id FROM calendars
			WHERE calendars.user_id  = ?
			AND calendars.checkout_time IS NULL
			AND DATEDIFF(calendars.checkin_time, ?) = 0`, userID, timeNow).Scan(&id)

	if id == 0 {
		err = modelErrors.NewAppError(errors.New("you have not checkin or you have already checkout"), modelErrors.ValidationError)
	}

	return
}

func Checkin(cal *Calendar) (err error) {

	if err := CheckValidCheckin(cal.UserID); err != nil {
		return modelErrors.NewAppError(err, modelErrors.ValidationError)
	}

	err = db.DB.Create(cal).Error

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

func Checkout(id int, calMap map[string]interface{}) (cal Calendar, err error) {
	cal.ID = id
	err = db.DB.Model(&cal).
		Select("checkin_time", "checkout_time").
		Updates(calMap).
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

	err = db.DB.Where("id = ?", id).First(&cal).Error

	if err != nil {
		err = modelErrors.NewAppErrorWithType(modelErrors.NotFound)
		return
	}

	return
}
