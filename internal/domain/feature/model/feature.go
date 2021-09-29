package model

type Feature struct {
	ID            string
	DisplayName   string
	TechnicalName string
	ExpiresOn     int64
	Description   string
	Inverted      bool
	CustomerIDs   []string
	CreatedAt     int64
	UpdatedAt     int64
	DeletedAt     int64
}
