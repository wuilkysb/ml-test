package router

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestRouterSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}

type RouterTestSuite struct {
	suite.Suite

	server    *echo.Echo
	underTest *Router
}

func (suite *RouterTestSuite) SetupTest() {
	suite.server = echo.New()
	suite.underTest = NewRouter(suite.server)
}

func (suite *RouterTestSuite) TestInit() {
	//apiGroup := suite.server.Group(enums.BasePath)

	suite.underTest.Init()
}
