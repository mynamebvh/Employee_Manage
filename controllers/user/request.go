package user

import "time"

type NewUser struct {
	EmployeeCode string    `json:"employee_code" example:"001" validate:"required"`
	FullName     string    `json:"full_name" example:"Bui Viet Hoang" validate:"required"`
	Password     string    `json:"password" example:"hoangdz" validate:"required"`
	Phone        string    `json:"phone" example:"0979150931" validate:"required"`
	Email        string    `json:"email" example:"mynamebvh@gmail.com" validate:"required"`
	Gender       bool      `json:"gender" example:"true" validate:"required"`
	Birthday     time.Time `json:"birthday" example:"2022-10-07T08:43:38+00:00" validate:"required"`
	Address      string    `json:"address" example:"Hoai Duc, Ha Noi" validate:"required"`
	DepartmentID int       `json:"department_id" example:"1" validate:"required"`
	RoleID       int       `json:"role_id" example:"1" validate:"required"`
}

type RequestChangePassword struct {
	OldPassword string `json:"old_password" example:"hoangdz" validate:"required"`
	NewPassword string `json:"new_password" example:"hoangdz1" validate:"required"`
}
