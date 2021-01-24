package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Health struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

func HealthCheck(c echo.Context) error {
	response := &Health{
		Code:    http.StatusOK,
		Message: "Active!",
	}

	return c.JSON(http.StatusOK, response)
}
