package models

import (
	"time"
)

type User struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	EmployeeCode string    `json:"employee_code"`
	FullName     string    `json:"full_name"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Gender       bool      `json:"gender"`
	Birthday     time.Time `json:"birthday"`
	Address      string    `json:"address"`
	Role         string    `json:"role"`
	Password     string    `json:"password"`
	DepartmentID int       `gorm:"not null;column:department_id" json:"department_id"`
	Department   Department
	CreatedAt    time.Time `json:"created_at,omitempty" gorm:"autoCreateTime:mili"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime:mili"`
}
