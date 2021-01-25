package controllers

import (
	"github.com/labstack/echo/v4"
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

func HealthCheck(c echo.Context) error {
	response := &Health{
		Code:    http.StatusOK,
		Message: "Active!",
	}

	return c.JSON(http.StatusOK, response)
}
