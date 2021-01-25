package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"ml-mutant-test/db/models"
	"ml-mutant-test/interfaces/controller"
	"ml-mutant-test/mocks"
	"net/http"
	"testing"
)

var (
	someError = fmt.Errorf("some error")
	mutantRequestFail = models.Mutant{
		DNA:      []string{"aae", "aaa"},
	}
	mutantRequest = models.Mutant{
		DNA:      []string{"aaa", "aaa", "aaa"},
	}
)

func TestMutantControllerSuite(t *testing.T) {
	suite.Run(t, new(MutantControllerTestSuite))
}

type MutantControllerTestSuite struct {
	suite.Suite
	mutantService *mocks.MutantServiceInterface
	underTest      controller.MutantControllerInterface
}

func (suite *MutantControllerTestSuite) SetupTest() {
	suite.mutantService = &mocks.MutantServiceInterface{}
	suite.underTest = NewMutantController(suite.mutantService)
}

func (suite *MutantControllerTestSuite) TestStats_WhenSuccess() {
	c := SetupControllerCase(http.MethodGet, "/stats", nil)

	suite.mutantService.Mock.On("Stats").Return(models.Stats{}, nil)

	err := suite.underTest.Stats(c.context)

	suite.NoError(err)
	suite.Equal(http.StatusOK, c.Res.Code)
}

func (suite *MutantControllerTestSuite) TestStats_WhenServiceFail() {
	c := SetupControllerCase(http.MethodGet, "/stats", nil)

	suite.mutantService.Mock.On("Stats").Return(models.Stats{}, someError)

	err := suite.underTest.Stats(c.context)

	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, err.(*echo.HTTPError).Code)
}

func (suite *MutantControllerTestSuite) TestIsMutant_WhenFailByBind() {
	body, _ := json.Marshal("")
	c := SetupControllerCase(http.MethodPost, "/mutant/", bytes.NewBuffer(body))
	c.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	err := suite.underTest.IsMutant(c.context)

	suite.Error(err)
	suite.Equal(http.StatusBadRequest, err.(*echo.HTTPError).Code)
}

func (suite *MutantControllerTestSuite) TestIsMutant_WhenFailByValidator() {
	body, _ := json.Marshal(mutantRequestFail)
	c := SetupControllerCase(http.MethodPost, "/mutant/", bytes.NewBuffer(body))
	c.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	err := suite.underTest.IsMutant(c.context)

	suite.Error(err)
	suite.Equal(http.StatusBadRequest, err.(*echo.HTTPError).Code)
}

func (suite *MutantControllerTestSuite) TestIsMutant_WhenServiceFail() {
	body, _ := json.Marshal(mutantRequest)
	c := SetupControllerCase(http.MethodPost, "/mutant/", bytes.NewBuffer(body))
	c.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.mutantService.Mock.On("IsMutant", mutantRequest.DNA).Return(false)

	err := suite.underTest.IsMutant(c.context)

	suite.Error(err)
	suite.Equal(http.StatusForbidden, err.(*echo.HTTPError).Code)
}

func (suite *MutantControllerTestSuite) TestIsMutant_WhenSuccess() {
	body, _ := json.Marshal(mutantRequest)
	c := SetupControllerCase(http.MethodPost, "/mutant/", bytes.NewBuffer(body))
	c.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.mutantService.Mock.On("IsMutant", mutantRequest.DNA).Return(true)

	err := suite.underTest.IsMutant(c.context)

	suite.NoError(err)
	suite.Equal(http.StatusOK, c.Res.Code)
}