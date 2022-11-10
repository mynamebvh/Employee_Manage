package mocks

import (
	. "employee_manage/tests/config"

	"github.com/DATA-DOG/go-sqlmock"
)

func MockGetDepartments() {
	MockAdmin()
	rows := Mock.NewRows([]string{"id", "name"}).
		AddRow(1, "Phòng 1")
	rowCount := Mock.NewRows([]string{"count"}).
		AddRow(1)
	Mock.ExpectQuery("SELECT (.+) FROM `departments`").WillReturnRows(rows)
	Mock.ExpectQuery("SELECT count(.+)").WillReturnRows(rowCount)
}

func MockGetDepartmentByID() {
	MockAdmin()
	rows := Mock.NewRows([]string{"id", "name"}).
		AddRow(1, "Phòng 1")
	Mock.ExpectQuery("SELECT (.+) FROM `departments` (.+)").WillReturnRows(rows)
}

func MockFailGetDepartmentByID() {
	MockAdmin()
	rows := Mock.NewRows([]string{"id", "name"})
	Mock.ExpectQuery("SELECT (.+) FROM `departments` (.+)").WillReturnRows(rows)
}

func MockCreateDepartment() {
	MockAdmin()
	Mock.ExpectBegin()
	Mock.ExpectExec("INSERT INTO `departments` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	Mock.ExpectCommit()
}

func MockUpdateDepartmentByID() {
	MockAdmin()
	rows := Mock.NewRows([]string{"id", "name"}).
		AddRow(4, "Phòng 4")
	Mock.ExpectQuery("SELECT (.+) FROM `departments` (.+)").WillReturnRows(rows)
}

func MockDeleteDepartmentByID() {
	MockAdmin()
	Mock.ExpectBegin()
	Mock.ExpectExec("DELETE FROM `departments` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	Mock.ExpectCommit()
}

func MockFailDeleteDepartmentByID() {
	MockAdmin()
	Mock.ExpectBegin()
	Mock.ExpectExec("DELETE FROM `departments` (.+)").WillReturnResult(sqlmock.NewResult(0, 0))
	Mock.ExpectCommit()
}
