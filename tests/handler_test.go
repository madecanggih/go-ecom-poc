package tests

import (
	"architect/saras-go-poc/handlers"
	"architect/saras-go-poc/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type (
	UserModelStub struct{}
)

func (u *UserModelStub) FindByID(id int) models.Users {
	return models.Users{
		ID:   1,
		Name: "foo",
	}
}

func (u *UserModelStub) FindAll() []models.Users {
	users := []models.Users{}
	users = append(users, models.Users{
		ID:   100,
		Name: "foo",
	})
	return users
}

func TestGetIndex(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.GET, "/index", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	u := &UserModelStub{}
	h := handlers.NewHandler(u)

	var expected = `{"users":[{"id": 100, "name": "foo"}]}`

	if assert.NoError(t, h.GetIndex(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
