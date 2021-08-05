package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/bensaufley/aud-it/internal/db"
	"github.com/bensaufley/aud-it/internal/graphql"
	"github.com/bensaufley/aud-it/internal/schema"
	"github.com/bensaufley/aud-it/internal/server"
)

func main() {
	cfg := &graphql.Config{
		DB: &db.Config{
			DBPath:         "/storage/data.db",
			MigrationsPath: "migrations",
		},
		SchemaString: schema.String,
	}

	mux, err := server.New(cfg)
	if err != nil {
		log.WithError(err).Fatal("could noot initialize server")
	}

	log.Info("server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.WithError(err).Fatal("could not start server")
	}
}
