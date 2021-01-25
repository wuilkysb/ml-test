package router

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"ml-mutant-test/mocks"
	"testing"
)

func TestRouterSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}

type RouterTestSuite struct {
	suite.Suite

	server           *echo.Echo
	mutantController *mocks.MutantControllerInterface
	underTest        *Router
}

func (suite *RouterTestSuite) SetupTest() {
	suite.server = echo.New()
	suite.mutantController = &mocks.MutantControllerInterface{}
	suite.underTest = NewRouter(suite.server, suite.mutantController)
}

func (suite *RouterTestSuite) TestInit() {
	//apiGroup := suite.server.Group(enums.BasePath)

	suite.underTest.Init()
}
