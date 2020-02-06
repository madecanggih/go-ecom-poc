package helpers

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func MockDB(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	_, mock, err := sqlmock.NewWithDSN("sqlmock_db_0")
	if err != nil {
		t.Fatalf("An error '%s' was not expecting", err)
	}

	db, err := gorm.Open("sqlmock", "sqlmock_db_0")
	return mock, db
}
