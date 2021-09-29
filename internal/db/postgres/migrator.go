package postgres

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func Migrate(db *sqlx.DB, path, database string) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to create postgres driver")
	}
	migration, err := migrate.NewWithDatabaseInstance(path, database, driver)
	if err != nil {
		return errors.Wrapf(err, "failed to connect to postgres for migration; path %q, database %q", path, database)
	}
	if err = migration.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return errors.Wrapf(err, "failed to migrate postgres; path %q, database %q", path, database)
	}
	return nil
}
