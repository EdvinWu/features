package entity

import (
	"database/sql"
	"time"
)

type Feature struct {
	ID            string       `db:"id"`
	DisplayName   string       `db:"display_name"`
	TechnicalName string       `db:"technical_name"`
	ExpiresOn     sql.NullTime `db:"expires_on"`
	Description   string       `db:"description"`
	Inverted      bool         `db:"inverted"`
	CustomerIDs   string       `db:"customer_ids"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
	DeletedAt     sql.NullTime `db:"deleted_at"`
}
