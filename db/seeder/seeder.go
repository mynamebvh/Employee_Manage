package main

import (
	"employee_manage/config"
	"employee_manage/models"

	"fmt"
	"time"
)

func main() {
	config.LoadEnv()

	birthday, _ := time.Parse("2006-01-02 15:04:05.000 -0700", "2020-01-02 03:04:05.000 +0000")

	users := []models.User{
		{ID: 1, FullName: "Root", Phone: "0979150931", Email: "mynamebvh@gmail.com", Gender: true, Address: "HN", Password: "hoangdz", RoleID: 1, DepartmentID: 1, Birthday: birthday},
		{ID: 2, FullName: "Manager D1", Phone: "0979150932", Email: "mynamebvh1@gmail.com", Gender: true, Address: "HN", Password: "hoangdz", RoleID: 2, DepartmentID: 1, Birthday: birthday},
		{ID: 3, FullName: "Bui Viet Hoang", Phone: "0979150933", Email: "mynamebvh2@gmail.com", Gender: true, Address: "HN", Password: "hoangdz", RoleID: 3, DepartmentID: 1, Birthday: birthday},
	}

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

	userDepartments := []models.UserDepartment{
		{UserID: 2, DepartmentID: 1},
	}

	db, err := config.GormOpen()

	if err != nil {
		fmt.Printf("Error connect database %s", err.Error())
	}

	fmt.Println("========Connect database successful========")
	time.Sleep(2 * time.Second)

	db.Create(&roles)
	db.Create(&departments)
	db.Create(&users)
	db.Create(&userDepartments)

	fmt.Println("========Seeder successful========")
}
