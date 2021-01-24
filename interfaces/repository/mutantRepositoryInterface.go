package repository

import "ml-mutant-test/db/models"

type MutantRepositoryInterface interface {
	Create(mutant *models.Mutant)
	GetByDNA(dna []string) (models.Mutant, error)
	GetStats() models.Stats
}
