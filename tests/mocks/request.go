package mocks

import (
	. "employee_manage/tests/config"

	"github.com/DATA-DOG/go-sqlmock"
)

func MockGetRequests() {
	MockAdmin()
	rows := Mock.NewRows([]string{"id", "type", "content"}).
		AddRow(1, "take_leave", "em đưa vợ đi đẻ")
	rowCount := Mock.NewRows([]string{"count"}).
		AddRow(1)
	Mock.ExpectQuery("SELECT (.+) FROM requests (.+)").WillReturnRows(rows)
	Mock.ExpectQuery("SELECT count(.+)").WillReturnRows(rowCount)
}

func MockGetRequestByID() {
	MockAdmin()
	row := Mock.NewRows([]string{"id", "type", "content"}).
		AddRow(1, "take_leave", "em đưa vợ đi đẻ")
	Mock.ExpectQuery("SELECT (.+) FROM requests (.+)").WillReturnRows(row)
}

func MockFailGetRequestByID() {
	MockAdmin()
	row := Mock.NewRows([]string{"id", "type", "content"})
	Mock.ExpectQuery("SELECT (.+) FROM requests (.+)").WillReturnRows(row)
}

func MockCreateRequest() {
	Mock.ExpectBegin()
	Mock.ExpectExec("INSERT INTO `requests` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	Mock.ExpectCommit()
}

func MockUpdateRequestByID() {
	MockAdmin()
	row := Mock.NewRows([]string{"id", "type", "content"}).
		AddRow(1, "take_leave", "em đưa vợ đi đẻ")
	Mock.ExpectQuery("SELECT (.+) FROM `requests` (.+)").WillReturnRows(row)
}
