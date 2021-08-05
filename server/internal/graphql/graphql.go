package graphql

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	log "github.com/sirupsen/logrus"

	"github.com/bensaufley/aud-it/internal/db"
	"github.com/bensaufley/aud-it/internal/resolver"
)

type Config struct {
	DB           *db.Config
	SchemaString func() (string, error)
}

func (cfg *Config) NewHandler() (*relay.Handler, error) {
	s, err := cfg.SchemaString()
	if err != nil {
		log.WithError(err).Fatal("could not build schema string")
	}
	sqlite, err := cfg.DB.Get()
	if err != nil {
		log.WithError(err).Fatal("error initializing database")
	}
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schm, err := graphql.ParseSchema(s, resolver.NewRoot(sqlite), opts...)
	if err != nil {
		log.WithError(err).Fatal("could not parse schema")
	}
	return &relay.Handler{Schema: schm}, nil
}
