package controllers

import "time"

type NewUser struct {
	EmployeeCode string    `json:"employee_code" example:"001" validate:"required"`
	FullName     string    `json:"full_name" example:"Bui Viet Hoang" validate:"required"`
	Phone        string    `json:"phone" example:"0979150931" validate:"required"`
	Email        string    `json:"email" example:"mynamebvh@gmail.com" validate:"required"`
	Gender       bool      `json:"gender" example:"true" validate:"required"`
	Birthday     time.Time `json:"birthday" example:"2022-10-07T08:43:38+00:00" validate:"required"`
	Address      string    `json:"address" example:"Hoai Duc, Ha Noi" validate:"required"`
	DepartmentID int       `json:"department_id" example:"1" validate:"required"`
	RoleID       int       `json:"role_id" example:"1" validate:"required"`
}
