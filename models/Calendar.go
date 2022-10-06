package models

import "time"

type Calendar struct {
	Type         string `json:"type"`
	UserID       int    `gorm:"not null;column:user_id" json:"user_id"`
	User         User
	CheckinTime  time.Time `json:"checkin_time,omitempty" gorm:"autoCreateTime:mili"`
	CheckoutTime time.Time `json:"checkout_time"`
}
