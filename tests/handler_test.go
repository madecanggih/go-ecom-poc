package tests

import (
	"architect/saras-go-poc/handlers"
	"architect/saras-go-poc/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedObject struct {
	mock.Mock
}

func (m *MockedObject) GetUsers(id string) []models.Users {
	users := []models.Users{}
	return users
}

func TestGetUser(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.GET, "/api/v1/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
