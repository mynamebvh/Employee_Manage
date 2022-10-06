package models

type Report struct {
	Type       string `json:"type"`
	Content    string `json:"content"`
	Status     bool   `json:"status"`
	UserID     int    `gorm:"not null;column:user_id" json:"user_id"`
	User       User
	ApprovedBy int  `gorm:"not null;column:approved_by" json:"approved_by"`
	Manager    User `gorm:"foreignKey:ApprovedBy"`
}
