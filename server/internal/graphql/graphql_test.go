package graphql_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bensaufley/graphql-preact-starter/internal/db"
	"github.com/bensaufley/graphql-preact-starter/internal/graphql"
	"github.com/bensaufley/graphql-preact-starter/internal/schema"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	dir := t.TempDir()
	tmpFile := filepath.Clean(dir + "/test.db")
	_, err := os.Create(tmpFile)
	if err != nil {
		assert.FailNow(t, "could not create temp database file", err)
	}

	cfg := graphql.Config{
		DB: &db.Config{
			DBPath: tmpFile,
		},
		SchemaString: schema.String,
	}
	_, err = cfg.NewHandler()

	assert.NoError(t, err)
}
