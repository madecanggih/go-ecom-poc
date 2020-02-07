package tests

import (
	"testing"

	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
)

func MockDB(t *testing.T) *gorm.DB {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	db, _ := gorm.Open(mocket.DriverName, "mock_db")

	return db
}
