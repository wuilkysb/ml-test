package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"ml-mutant-test/enums"
	"net/http"
	"net/http/httptest"
)

type Health struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

type ControllerCase struct {
	Req     *http.Request
	Res     *httptest.ResponseRecorder
	context echo.Context
}

func SetupControllerCase(method string, url string, body io.Reader) ControllerCase {
	path := fmt.Sprintf(enums.BasePath+"%s", url)

	e := echo.New()
	req := httptest.NewRequest(method, path, body)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	return ControllerCase{req, res, c}
}

func HealthCheck(c echo.Context) error {
	response := &Health{
		Code:    http.StatusOK,
		Message: "Active!",
	}

	return c.JSON(http.StatusOK, response)
}
