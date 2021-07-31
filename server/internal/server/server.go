package server

import (
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"

	"github.com/bensaufley/graphql-preact-starter/internal/graphiql"
	"github.com/bensaufley/graphql-preact-starter/internal/graphql"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

var fs = http.Dir("/public")

func HandleStatic(extPath string, intPath string) http.Handler {
	publicRegExp := regexp.MustCompile(fmt.Sprintf("^%s", intPath))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, err := filepath.Rel(extPath, r.URL.Path)
		if err != nil {
			log.WithError(err).Warn("could not get relative path from static")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		pp := filepath.Clean(fmt.Sprintf("%s/%s", intPath, p))
		if !publicRegExp.MatchString(pp) {
			log.WithField("path", pp).Warn("clean public filepath is not within public")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		http.ServeFile(w, r, pp)
	})
}

func New(cfg *graphql.Config) (*http.ServeMux, error) {
	gqlhandler, err := cfg.NewHandler()
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize GraphQL handler")
	}

	mux := http.NewServeMux()

	mux.Handle("/static/", HandleStatic("/static", "/public"))
	mux.Handle("/graphiql", graphiql.Serve)
	mux.Handle("/graphql", gqlhandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/public/index.html")
	})

	return mux, nil
}
