package mocks

import (
	. "employee_manage/tests/config"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func MockCheckIn() {
	rows := Mock.NewRows([]string{"COUNT(id)"}).
		AddRow(0)
	Mock.ExpectQuery("SELECT (.+) FROM calendars (.+)").WillReturnRows(rows)

	Mock.ExpectBegin()
	Mock.ExpectExec("INSERT INTO `calendars` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	Mock.ExpectCommit()
}

func MockCheckOut() {
	rowCheck := Mock.NewRows([]string{"id"}).
		AddRow(1)
	Mock.ExpectQuery("SELECT (.+) FROM calendars (.+)").WillReturnRows(rowCheck)

	rows := Mock.NewRows([]string{"id", "user_id", "checkin_time", "checkout_time"}).
		AddRow(1, "3", time.Now(), time.Now())
	Mock.ExpectQuery("SELECT (.+) FROM `calendars` (.+)").WillReturnRows(rows)
}
