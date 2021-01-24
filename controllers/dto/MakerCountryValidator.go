package dto

type MakerCountryValidator struct {
	Maker             string `json:"maker"`
	ProductsTrademark string `json:"products_trademark"`
	Country           string `json:"country"`
	ProductID         string `json:"product_id"`
}
