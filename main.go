package main

import (
	"ml-mutant-test/config"
	"ml-mutant-test/db/migrations"
	"ml-mutant-test/providers"
	"ml-mutant-test/router"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/signalfx/signalfx-go-tracing/tracing"
)

func main() {
	tracing.Start()
	defer tracing.Stop()
	container := providers.BuildContainer()

	migrations.StartConfiguration()
	err := container.Invoke(func(server *echo.Echo, route *router.Router) {
		address := fmt.Sprintf("%s:%s", config.Environments().ServerHost, config.Environments().ServerPort)

		route.Init()
		server.Logger.Fatal(server.Start(address))
	})

	if err != nil {
		panic(err)
	}
}
