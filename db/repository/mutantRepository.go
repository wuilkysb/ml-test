package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/gommon/log"
	"ml-mutant-test/db/models"
	"ml-mutant-test/interfaces/repository"
)

type MutantRepository struct {
	db *pg.DB
}

func NewMutantRepository(
	db *pg.DB,
) repository.MutantRepositoryInterface {
	return &MutantRepository{
		db,
	}
}

func (r *MutantRepository) Create(mutant *models.Mutant) {
	if _, err := r.db.Model(mutant).Insert(); err != nil {
		log.Errorf("insert error: %s", err.Error())
	}
}

func (r *MutantRepository) GetByDNA(dna []string) (models.Mutant, error) {
	var mutant models.Mutant
	if err := r.db.Model(&mutant).Where("dna = ?", pg.Array(dna)).Select(); err != nil {
		log.Error("get error: %s", err.Error())
		return mutant, err
	}
	return mutant, nil
}

func (r *MutantRepository) GetStats() models.Stats {
	var stats models.Stats
	if err := r.db.Model().
		With("is_mutants", r.db.Model().TableExpr("mutants").ColumnExpr("count(id) as total").Where("is_mutant is true")).
		With("is_humans", r.db.Model().TableExpr("mutants").ColumnExpr("count(id) as total").Where("is_mutant is false")).
		TableExpr("is_mutants").
		ColumnExpr("is_humans.total as count_human_dna").
		ColumnExpr("is_mutants.total as count_mutant_dna").
		ColumnExpr("(is_mutants.total * 1.0/COALESCE(nullif(is_humans.total,0), 1))::float as ratio").
		Join("JOIN is_humans on true").
		Group("count_human_dna").Group("count_mutant_dna").Group("ratio").Select(&stats); err != nil {
			log.Infof("stats error: %s", err.Error())
			return stats
	}

	return stats
}
