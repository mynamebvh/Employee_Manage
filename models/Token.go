package models

import "time"

type Token struct {
	Value     string `json:"value"`
	Type      string `json:"type"`
	UserID    int    `gorm:"not null;column:user_id" json:"user_id"`
	User      User
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime:mili"`
}
