package department

type NewDepartment struct {
	Name           string `json:"name" example:"Phòng D5" validate:"required"`
	DepartmentCode string `json:"department_code" example:"D5" validate:"required"`
	Address        string `json:"address" example:"Tầng 3" validate:"required"`
	Status         bool   `json:"status" example:"true" validate:"required"`
}
