package services

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"ml-mutant-test/db/models"
	"ml-mutant-test/interfaces/services"
	"ml-mutant-test/mocks"
	"testing"
)

var (
	someError             = fmt.Errorf("some error")
	dnaWithWrongDimension = []string{"aaa", "aaa", "aaa", "aaa"}
	dna = []string{"aaa", "aaa", "aaa"}
	dnaUL = [][]string{
		{"a", "t", "g"},
		{"t", "a", "c"},
		{"t", "t", "a"},
	}
	dnaUR = [][]string{
		{"c", "t", "a"},
		{"t", "a", "c"},
		{"a", "t", "c"},
	}
	dnaR = [][]string{
		{"a", "a", "a"},
		{"t", "g", "c"},
		{"g", "t", "c"},
	}
	dnaU = [][]string{
		{"a", "t", "c"},
		{"g", "c", "c"},
		{"t", "a", "c"},
	}
	mutant = models.Mutant{
		DNA:      dna,
		IsMutant: true,
	}
)

type MutantServiceTestSuite struct {
	suite.Suite
	repo        *mocks.MutantRepositoryInterface
	underTest   services.MutantServiceInterface
}

func TestMutantServiceSuite(t *testing.T) {
	suite.Run(t, new(MutantServiceTestSuite))
}

func (suite *MutantServiceTestSuite) SetupTest() {
	suite.repo = &mocks.MutantRepositoryInterface{}
	suite.underTest = NewMutantService(suite.repo)
}

func (suite *MutantServiceTestSuite) TestStats_WhenRepoFail() {

	suite.repo.Mock.On("GetStats").Return(models.Stats{}, someError)
	_, err := suite.underTest.Stats()

	suite.Equal(someError.Error(), err.Error())
}

func (suite *MutantServiceTestSuite) TestStats_WhenSuccess() {

	suite.repo.Mock.On("GetStats").Return(models.Stats{}, nil)
	_, err := suite.underTest.Stats()

	suite.NoError(err)
}

func (suite *MutantServiceTestSuite) TestIsMutant_WhenFirstRepoCallGetResults() {
	suite.repo.Mock.On("GetByDNA", dna).Return(models.Mutant{}, nil)
	result := suite.underTest.IsMutant(dna)
	suite.False(result)
}

func (suite *MutantServiceTestSuite) TestIsMutant_WhenDNAHaveAWrongDimensions() {
	suite.repo.Mock.On("GetByDNA", dnaWithWrongDimension).Return(models.Mutant{}, someError)
	result := suite.underTest.IsMutant(dnaWithWrongDimension)
	suite.False(result)
}

func (suite *MutantServiceTestSuite) TestIsMutant_WhenSuccess() {
	suite.repo.Mock.On("GetByDNA", dna).Return(models.Mutant{}, someError)
	suite.repo.Mock.On("Create", &mutant)
	result := suite.underTest.IsMutant(dna)
	suite.True(result)
}

func (suite *MutantServiceTestSuite) TestFindSequence_GetTransversalUpToLeft() {
	result := FindSequence(dnaUL, 2, 2, 3, 3)
	suite.Equal(result, []string{"aaa"})
}

func (suite *MutantServiceTestSuite) TestFindSequence_GetTransversalUpToRight() {
	result := FindSequence(dnaUR, 0, 2, 3, 3)
	suite.Equal(result, []string{"aaa"})
}

func (suite *MutantServiceTestSuite) TestFindSequence_GetRow() {
	result := FindSequence(dnaR, 0, 0, 3, 3)
	suite.Equal(result, []string{"aaa"})
}

func (suite *MutantServiceTestSuite) TestFindSequence_GetColumn() {
	result := FindSequence(dnaU, 2, 2, 3, 3)
	suite.Equal(result, []string{"ccc"})
}