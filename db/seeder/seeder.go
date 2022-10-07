package main

import (
	"employee_manage/config"
	"employee_manage/models"

	"fmt"
	"time"
)

func main() {
	roles := []models.Role{
		{ID: 1, Name: "admin"},
		{ID: 2, Name: "manager"},
		{ID: 3, Name: "user"},
	}

	departments := []models.Department{
		{ID: 1, Name: "Phòng 1", DepartmentCode: "D1", Address: "Tầng 3A", Status: true},
		{ID: 2, Name: "Phòng 2", DepartmentCode: "D2", Address: "Tầng 3", Status: true},
		{ID: 3, Name: "Phòng 3", DepartmentCode: "D3", Address: "Tầng 3", Status: true},
	}

	db, err := config.GormOpen()

	if err != nil {
		fmt.Printf("Error migrate database %s", err.Error())
	}

	fmt.Println("========Connect database successful========")
	time.Sleep(2 * time.Second)

	db.Create(&roles)
	db.Create(&departments)

	fmt.Println("========Seeder successful========")
}
