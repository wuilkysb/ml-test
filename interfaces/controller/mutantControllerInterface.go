package controller

import "github.com/labstack/echo/v4"

type MutantControllerInterface interface {
	IsMutant(c echo.Context) error
	Stats(c echo.Context) error
}
