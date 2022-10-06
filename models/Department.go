package models

import (
	"time"
)

var (
	TEST_IMPORT = "TEST"
)

type Department struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	DepartmentCode string    `json:"department_code"`
	Address        string    `json:"address"`
	Status         bool      `json:"status"`
	CreatedAt      time.Time `json:"created_at,omitempty" gorm:"autoCreateTime:mili"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime:mili"`
}
