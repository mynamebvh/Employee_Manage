package controllers

import "time"

type NewUser struct {
	EmployeeCode string    `json:"employee_code"`
	FullName     string    `json:"full_name"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Gender       bool      `json:"gender"`
	Birthday     time.Time `json:"birthday"`
	Address      string    `json:"address"`
	DepartmentID int       `json:"department_id"`
	RoleID       int       `json:"role_id"`
}
