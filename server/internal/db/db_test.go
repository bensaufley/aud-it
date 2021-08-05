package db_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/bensaufley/aud-it/internal/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/stretchr/testify/assert"
)

func Test_Migrations(t *testing.T) {
	tmpDir := t.TempDir()
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s/test.db?mode=rwc", tmpDir))
	if err != nil {
		assert.FailNowf(t, err.Error(), "could not open database")
	}
	dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		assert.FailNowf(t, err.Error(), "could not initialize migration driver")
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"sqlite3",
		dbDriver,
	)
	if err != nil {
		assert.FailNowf(t, err.Error(), "could not initialize migrator")
	}
	if assert.NoError(t, m.Up()) {
		assert.NoError(t, m.Down())
	}
}

func TestGet(t *testing.T) {
	tmpDir := t.TempDir()
	dbConfig := &db.Config{
		DBPath:         fmt.Sprintf("%s/test.db", tmpDir),
		MigrationsPath: "../../migrations",
	}

	database, err := dbConfig.Get()

	if assert.NoError(t, err) {
		if assert.NotNil(t, database) {
			assert.IsType(t, &sql.DB{}, database)
		}
	}
}
