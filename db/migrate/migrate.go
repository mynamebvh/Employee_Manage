package main

import (
	"fmt"
	"time"

	"employee_manage/config"
	"employee_manage/models"
)

func main() {
	db, err := config.GormOpen()

	if err != nil {
		fmt.Printf("Error migrate database %s", err.Error())
	}

	fmt.Println("========Connect database successful========")
	time.Sleep(2 * time.Second)
	db.AutoMigrate(&models.Department{}, &models.Role{}, &models.User{}, &models.UserDepartment{}, &models.Request{}, &models.Token{}, &models.Calendar{})

	fmt.Println("========Migrate successful========")
}
