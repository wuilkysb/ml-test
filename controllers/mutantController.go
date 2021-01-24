package controllers

import (
	"github.com/labstack/echo/v4"
	"ml-mutant-test/db/models"
	"ml-mutant-test/interfaces/controller"
	"ml-mutant-test/interfaces/services"
	"net/http"
)

type MutantController struct {
	service services.MutantServiceInterface
}

func NewMutantController(service services.MutantServiceInterface) controller.MutantControllerInterface {
	return &MutantController{
		service,
	}
}

func (controller *MutantController) IsMutant(c echo.Context) error {
	mutant := &models.Mutant{}
	if err := c.Bind(&mutant); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "dna is required")
	}

	if err := mutant.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if controller.service.IsMutant(mutant.DNA) {
		return c.JSON(http.StatusOK, "ok")
	}
	return echo.NewHTTPError(http.StatusForbidden, "is not a mutant")
}

func (controller *MutantController) Stats(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.Stats())
}
