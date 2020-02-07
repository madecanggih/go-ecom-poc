package tests

import (
	"testing"

	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
)

func MockDB(t *testing.T) *gorm.DB { // or *gorm.DB
	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = true

	// GORM
	db, _ := gorm.Open(mocket.DriverName, "mock_db") // Can be any connection string

	return db
}
