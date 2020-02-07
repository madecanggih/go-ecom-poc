package tests

import (
	"architect/saras-go-poc/handlers"
	"architect/saras-go-poc/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type UserModelStub struct{}

func (u *UserModelStub) SelectById(id int) models.Users {
	return models.Users{
		ID:   1,
		Name: "foo",
	}
}

func (u *UserModelStub) SelectAll() []models.Users {
	users := []models.Users{
		models.Users{
			ID:   1,
			Name: "foo",
		},
		models.Users{
			ID:   2,
			Name: "bar",
		},
	}

	return users
}

func TestGetIndex(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.GET, "/index", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	u := &UserModelStub{}
	h := handlers.NewHandler(u)

	var expected = `[{"id":1,"email":"","username":"","name":"foo","address":"","phone":"","password":"","image":""},{"id":2,"email":"","username":"","name":"bar","address":"","phone":"","password":"","image":""}]`

	if assert.NoError(t, h.GetIndex(c)) {
		var actual = strings.TrimSpace(rec.Body.String())

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, actual)
	}
}

func TestGetDetail(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.GET, "/index/:1", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	u := &UserModelStub{}
	h := handlers.NewHandler(u)

	var expected = `{"id":1,"email":"","username":"","name":"foo","address":"","phone":"","password":"","image":""}`

	if assert.NoError(t, h.GetDetail(c)) {
		var actual = strings.TrimSpace(rec.Body.String())

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, actual)
	}
}
