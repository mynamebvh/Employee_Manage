package mocks

import (
	. "employee_manage/tests/config"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func MockGetUsers() {
	MockAdmin()

	rows := Mock.NewRows([]string{"id", "email"}).
		AddRow(1, "mynamebvh@gmail.com")
	rowCount := Mock.NewRows([]string{"count"}).
		AddRow(1)

	Mock.ExpectQuery("SELECT (.+) FROM `users` (.+)").WillReturnRows(rows)
	Mock.ExpectQuery("SELECT count(.+)").WillReturnRows(rowCount)
}

func MockGetUserByID() {
	MockAdmin()

	rows := Mock.NewRows([]string{"id", "email"}).
		AddRow(1, "mynamebvh@gmail.com")

	Mock.ExpectQuery("SELECT (.+) FROM `users` WHERE `users`.`id` = ?").WithArgs(3).WillReturnRows(rows)
}

func MockCreateUser() {
	MockAdmin()

	Mock.ExpectBegin()
	Mock.ExpectExec("INSERT INTO `users` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	Mock.ExpectCommit()
}

func MockRoleCreateUser() {
	MockUser()
}

func MockUpdateUserByID() {
	MockAdmin()

	rows := Mock.NewRows([]string{"id", "email"}).
		AddRow(4, "mynamebvh4@gmail.com")

	Mock.ExpectQuery("SELECT (.+) FROM `users` (.+)").WithArgs(4, 4).WillReturnRows(rows)

	// Mock.ExpectBegin()
	// Mock.ExpectExec("UPDATE users(.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	// Mock.ExpectCommit()
}

func MockDeleteUserByID() {
	MockAdmin()

	Mock.ExpectBegin()
	Mock.ExpectExec("DELETE FROM `users`").WithArgs(4).WillReturnResult(sqlmock.NewResult(1, 1))
	Mock.ExpectCommit()
}

func MockSendCode() {
	rows := Mock.NewRows([]string{"id", "email", "full_name"}).
		AddRow(5, "mynamebvh5@gmail.com", "Hoàng")

	Mock.ExpectQuery("SELECT (.+) FROM `users` WHERE email=?(.+)").WithArgs("mynamebvh5@gmail.com").WillReturnRows(rows)

	Mock.ExpectBegin()
	Mock.ExpectExec("INSERT INTO `tokens` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	Mock.ExpectCommit()
}

func MocKForgetPassword() {
	rows := Mock.NewRows([]string{"id", "email", "value", "expired"}).
		AddRow(1, "mynamebvh@gmail.com", "code", time.Now().Local().Add(time.Hour*time.Duration(2)))

	Mock.ExpectQuery("SELECT (.+) FROM `tokens` WHERE value=?").WithArgs("code").WillReturnRows(rows)

	Mock.ExpectBegin()
	Mock.ExpectExec("UPDATE `users` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	Mock.ExpectCommit()

	Mock.ExpectBegin()
	Mock.ExpectExec("DELETE `tokens` WHERE value = ?").WithArgs("code").WillReturnResult(sqlmock.NewResult(1, 1))
	Mock.ExpectCommit()
}
