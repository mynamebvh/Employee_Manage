package models

import "time"

type Calendar struct {
	ID           int `json:"id" gorm:"primaryKey"`
	UserID       int `gorm:"not null;column:user_id" json:"user_id"`
	User         User
	CheckinTime  time.Time `json:"checkin_time,omitempty" gorm:"autoCreateTime:mili"`
	CheckoutTime time.Time `json:"checkout_time"`
}
