package tests

import (
	"architect/saras-go-poc/helpers"
	"architect/saras-go-poc/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFindByID(t *testing.T) {
	mock, db := helpers.MockDB(t)
	defer db.Close()

	var cols []string = []string{"id", "name"}
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).AddRow(1, "foobar"))

	um := models.NewDB(db)
	u := um.FindByID(1)

	expect := models.Users{
		ID:   1,
		Name: "foobar",
	}
	assert.Equal(t, expect, u)
}
