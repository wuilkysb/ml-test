package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"regexp"
)

const (
	DNAStringRegex = "^[atgcATGC]+$"
)

var (
	DNARegex = regexp.MustCompile(DNAStringRegex)
)

type Mutant struct {
	DNA      []string `json:"dna" pg:",array" validate:"required,min=3,dive,dna_valid_characters,required"`
	IsMutant bool     `json:"is_mutant"`
}

func (u *Mutant) Validate() error {
	validate := validator.New()
	registerCustomDNAValidator(validate)
	return validate.Struct(u)
}

func registerCustomDNAValidator(v *validator.Validate) {
	if err := v.RegisterValidation("dna_valid_characters", IsValidSequence); err != nil {
		log.Error(err)
	}
}

func IsValidSequence(fl validator.FieldLevel) bool {
	return DNARegex.MatchString(fl.Field().String())
}
