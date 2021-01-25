package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"ml-mutant-test/enums"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var healthJson = `{"status":200,"message":"Active!"}`

func TestHealthCheck(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/one/catalog/management/health", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	if assert.NoError(t, HealthCheck(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, healthJson, strings.TrimSpace(res.Body.String()))
	}
}

func SetupControllerCase(method string, url string, body io.Reader) ControllerCase {
	path := fmt.Sprintf(enums.BasePath+"%s", url)

	e := echo.New()
	req := httptest.NewRequest(method, path, body)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	return ControllerCase{req, res, c}
}
