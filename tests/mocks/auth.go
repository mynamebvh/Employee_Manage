package mocks

import (
	// . "employee_manage/tests/config/config"

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
