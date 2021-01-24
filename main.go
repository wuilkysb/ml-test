package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/signalfx/signalfx-go-tracing/tracing"
	"ml-mutant-test/config"
	"ml-mutant-test/db/migrations"
	"ml-mutant-test/providers"
	"ml-mutant-test/router"
)

func main() {
	tracing.Start()
	defer tracing.Stop()
	container := providers.BuildContainer()

	migrations.StartConfiguration()
	err := container.Invoke(func(server *echo.Echo, route *router.Router) {
		address := fmt.Sprintf(":%s", config.Environments().Port)

		route.Init()
		server.Logger.Fatal(server.Start(address))
	})

	if err != nil {
		panic(err)
	}
}
