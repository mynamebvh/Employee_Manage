package dto

import "time"

type QueryResultGetMe struct {
	ID             int       `json:"user_id" `
	EmployeeCode   string    `json:"employee_code"`
	FullName       string    `json:"full_name"`
	DepartmentName string    `json:"department_name"`
	Phone          string    `json:"phone" `
	Email          string    `json:"email" `
	Gender         bool      `json:"gender" `
	Birthday       time.Time `json:"birthday"`
	Address        string    `json:"address"`
}
