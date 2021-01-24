package dto

type CatalogManagement struct {
	Maker     string   `json:"maker"`
	Type      string   `json:"type"`
	Catalogs  []string `json:"catalogs"`
	TableName string   `json:"table_name"`
}

type StatusManagement struct {
	Approval      string `json:"approval"`
	ApprovalEmail string `json:"approval_email"`
	Status        string `json:"status"`
	NewUrlFile    string `json:"url_file"`
}

type Uploader struct {
	UUID  string `json:"uuid"`
	Email string `json:"email"`
}
