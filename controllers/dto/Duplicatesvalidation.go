package dto

type DuplicateValidation struct {
	ProductID   string `json:"product_id"`
	Country     string `json:"country"`
	Occurrences int    `json:"occurrences"`
}
