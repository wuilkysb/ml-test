package providers

import (
	"go.uber.org/dig"
	"ml-mutant-test/controllers"
	database "ml-mutant-test/db"
	"ml-mutant-test/db/repository"
	"ml-mutant-test/router"
	"ml-mutant-test/server"
	"ml-mutant-test/services"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(server.NewServer)
	_ = Container.Provide(database.ConnInstance)

	_ = Container.Provide(router.NewRouter)

	_ = Container.Provide(controllers.NewMutantController)
	_ = Container.Provide(services.NewMutantService)

	_ = Container.Provide(repository.NewMutantRepository)

	return Container
}
