package postgres

import (
	"features/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func Connect(cfg *config.Postgres, connectionString string) (*sqlx.DB, error) {
	c, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, errors.New("failed to connect to postgres")
	}

	c.DB.SetMaxIdleConns(cfg.MaxIdleConnections)
	c.DB.SetMaxOpenConns(cfg.MaxOpenConnections)
	c.DB.SetConnMaxLifetime(cfg.MaxConnectionTTL)

	return c, nil
}

func URL(settings *config.Postgres) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		settings.Username,
		settings.Password,
		settings.Host,
		settings.Port,
		settings.Database)
}
