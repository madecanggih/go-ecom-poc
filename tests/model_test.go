package tests

import (
	"architect/saras-go-poc/models"
	"testing"

	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
)

func TestSelectById(t *testing.T) {
	db := MockDB(t)
	defer db.Close()

	mockReply := []map[string]interface{}{{"id": 1, "name": "foobar"}}
	mocket.Catcher.Reset().NewMock().WithQuery(`SELECT * FROM "users"  WHERE ("users"."id" = 1)`).WithReply(mockReply)

	um := models.NewDB(db)
	u := um.SelectById(1)

	expect := models.Users{
		ID:   1,
		Name: "foobar",
	}
	assert.Equal(t, expect, u)
}

func TestSelectAll(t *testing.T) {
	db := MockDB(t)
	defer db.Close()

	mockReply := []map[string]interface{}{{"id": 1, "name": "foo"}, {"id": 2, "name": "bar"}}
	mocket.Catcher.Reset().NewMock().WithQuery(`SELECT * FROM "users"`).WithReply(mockReply)

	um := models.NewDB(db)
	u := um.SelectAll()

	expect := []models.Users{
		models.Users{
			ID:   1,
			Name: "foo",
		},
		models.Users{
			ID:   2,
			Name: "bar",
		},
	}
	assert.Equal(t, expect, u)
}
