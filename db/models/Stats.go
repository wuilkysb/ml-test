package models

type Stats struct {
	CountHumanDNA  int     `json:"count_human_dna"`
	CountMutantDNA int     `json:"count_mutant_dna"`
	Ratio          float32 `json:"ratio"`
}
