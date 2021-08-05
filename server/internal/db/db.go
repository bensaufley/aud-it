package db

import (
	"database/sql"
	_ "embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type Config struct {
	DBPath         string
	MigrationsPath string
}

func (cfg *Config) Get() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=rwc", cfg.DBPath))
	if err != nil {
		return nil, errors.Wrap(err, "could not open database")
	}
	dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize migration driver")
	}
	if cfg.MigrationsPath != "" {
		m, err := migrate.NewWithDatabaseInstance(
			fmt.Sprintf("file://%s", cfg.MigrationsPath),
			"sqlite3",
			dbDriver,
		)
		if err != nil {
			return nil, errors.Wrap(err, "could not initialize migrator")
		}
		err = m.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return nil, errors.Wrap(err, "could not run migrations")
		}
	}
	return db, nil
}
