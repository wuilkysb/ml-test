package services

import "ml-mutant-test/db/models"

type MutantServiceInterface interface {
	IsMutant(dna []string) bool
	Stats() (models.Stats, error)
}
