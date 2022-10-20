package dto

type QueryGetWorkingTimeInMonth struct {
	UserID         int    `json:"user_id"`
	FullName       string `json:"full_name"`
	EmployeeCode   string `json:"employee_code"`
	Days           int    `json:"days"`
	Hours          string `json:"hours"`
	DepartmentName string `json:"department_name"`
}
