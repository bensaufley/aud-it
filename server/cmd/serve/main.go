package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/bensaufley/graphql-preact-starter/internal/db"
	"github.com/bensaufley/graphql-preact-starter/internal/graphql"
	"github.com/bensaufley/graphql-preact-starter/internal/schema"
	"github.com/bensaufley/graphql-preact-starter/internal/server"
)

func main() {
	cfg := &graphql.Config{
		DB: &db.Config{
			DBPath:         "/storage/data.db",
			MigrationsPath: "migrations",
		},
		SchemaString: schema.String,
	}

	if err := cfg.DB.InitFile(); err != nil {
		log.WithError(err).Fatal("could not initialize database")
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
