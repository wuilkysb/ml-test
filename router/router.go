package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ml-mutant-test/controllers"
	"ml-mutant-test/enums"
	"ml-mutant-test/interfaces/controller"
)

type Router struct {
	server           *echo.Echo
	mutantController controller.MutantControllerInterface
}

func NewRouter(
	server *echo.Echo,
	mutantController controller.MutantControllerInterface) *Router {
	return &Router{
		server,
		mutantController,
	}
}

func (r *Router) Init() {
	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))
	apiGroup := r.server.Group(enums.BasePath)

	apiGroup.GET(enums.Health, controllers.HealthCheck)
	apiGroup.POST("/", r.mutantController.IsMutant)
	apiGroup.GET("/stats", r.mutantController.Stats)
}
