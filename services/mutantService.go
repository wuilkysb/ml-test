package services

import (
	"fmt"
	"github.com/thoas/go-funk"
	"ml-mutant-test/db/models"
	"ml-mutant-test/interfaces/repository"
	"ml-mutant-test/interfaces/services"
	"strings"
)

const SequenceLength = 3

var (
	SequenceSearch = []string{"aaa", "ttt", "ccc", "ggg"}
)

type mutantService struct {
	repository repository.MutantRepositoryInterface
}

func NewMutantService(repository repository.MutantRepositoryInterface) services.MutantServiceInterface {
	return &mutantService{
		repository,
	}
}


func (s *mutantService) Stats() (models.Stats, error) {
	return s.repository.GetStats()
}

func (s *mutantService) IsMutant(dna []string) bool {
	if mutant, err := s.repository.GetByDNA(dna); err == nil {
		return mutant.IsMutant
	}
	var dnaInput [][]string
	N := len(dna)
	for i := range dna {
		if N != len(dna[i]) {
			return false
		}
		dnaInput = append(dnaInput, strings.Split(strings.ToLower(dna[i]), ""))
	}
	isMutant := sequencesSearch(dnaInput)
	mutant := &models.Mutant{
		DNA:      dna,
		IsMutant: isMutant,
	}
	s.repository.Create(mutant)
	return isMutant
}

func sequencesSearch(puzzle [][]string) bool {
	result := make([]string, 0, 0)
	N := len(puzzle)
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			w := FindSequence(puzzle, col, row, N, N)
			if len(w) > 0 {
				result = append(result, w...)
			}

		}
	}
	return len(result) > 1
}

func FindSequence(puzzle [][]string, xPos int, yPos int, xSize int, ySize int) []string {
	var resultUL, resultU, resultUR, resultR string
	for i := 0; i < SequenceLength; i++ {
		if (0 <= (xPos-SequenceLength+1) && xPos <= xSize) && (0 <= (yPos-SequenceLength+1) && yPos <= ySize) {
			resultUL += puzzle[yPos-i][xPos-i]
		}
		if (0 <= (yPos-SequenceLength+1) && yPos <= ySize) && (0 <= xPos && (xPos+SequenceLength) <= xSize) {
			resultUR += puzzle[yPos-i][xPos+i]
		}
		if 0 <= (yPos-SequenceLength+1) && yPos <= ySize {
			resultU += puzzle[yPos-i][xPos]
			fmt.Println(yPos)
		}
		if 0 <= xPos && (xPos+SequenceLength) <= xSize {
			resultR += puzzle[yPos][xPos+i]
		}
	}

	return ContentSequence([]string{resultUL, resultU, resultUR, resultR})
}

func ContentSequence(sequences []string) []string {
	return funk.IntersectString(SequenceSearch, sequences)
}
