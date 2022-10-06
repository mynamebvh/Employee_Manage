package models

type Manager struct {
	UserID       int `gorm:"primaryKey;not null;column:user_id" json:"user_id"`
	User         User
	DepartmentID int `gorm:"primaryKey;not null;column:department_id" json:"department_id"`
	Department   Department
}
