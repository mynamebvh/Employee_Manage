package mocks

import (
	. "employee_manage/tests/config"

	"github.com/DATA-DOG/go-sqlmock"
)

func MockAuth() {
	rows := Mock.NewRows([]string{"id", "email", "password"}).
		AddRow(1, "mynamebvh@gmail.com", "$2a$12$oGIJ9QJ43E6INO4usXcoGebRB7N1lfXKOVPpnj6.vzQcPnOm9SfDe")

	Mock.ExpectQuery("SELECT (.+) FROM `users` WHERE email=?(.+)").WithArgs("mynamebvh@gmail.com").WillReturnRows(rows)
	Mock.ExpectBegin()
	Mock.ExpectExec("INSERT INTO `tokens` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	Mock.ExpectCommit()
}

func MockAuthFail() {
	rows := Mock.NewRows([]string{"id", "email", "password"}).
		AddRow(1, "mynamebvh@gmail.com", "$2a$12$oGIJ9QJ43E6INO4usXcoGebRB7N1lfXKOVPpnj6.vzQcPnOm9SfDe")

	Mock.ExpectQuery("SELECT (.+) FROM `users` WHERE email=?(.+)").WithArgs("mynamebvh@gmail.com").WillReturnRows(rows)
}
func MockAdmin() {
	rowsAdmin := Mock.NewRows([]string{"department_id", "name"}).
		AddRow(nil, "admin")
	// Mock.ExpectQuery("SELECT (.+) FROM `users` left join user_departments (.+) WHERE users.id = ?").WithArgs(1).WillReturnRows(rowsAdmin)
	Mock.ExpectQuery("SELECT (.+) FROM `users` (.+)").WithArgs(1).WillReturnRows(rowsAdmin)
}

func MockManager() {
	rowsManager := Mock.NewRows([]string{"department_id", "name"}).
		AddRow(1, "manager")
	Mock.ExpectQuery("SELECT (.+) FROM `users` left join user_departments (.+) WHERE users.id = ?").WithArgs(2).WillReturnRows(rowsManager)
}

func MockUser() {
	rowsUser := Mock.NewRows([]string{"department_id", "name"}).
		AddRow(nil, "user")
	Mock.ExpectQuery("SELECT (.+) FROM `users` left join user_departments (.+) WHERE users.id = ?").WithArgs(3).WillReturnRows(rowsUser)
}
