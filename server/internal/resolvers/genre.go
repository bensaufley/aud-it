package resolvers

import (
	"context"
	"database/sql"

	"github.com/bensaufley/aud-it/internal/entities/genre"
)

type GenreResolver struct {
	db *sql.DB
	genre.Genre
}

func (r GenreResolver) Name(context.Context) (string, error) {
	return r.Genre.Name, nil
}

func (r GenreResolver) Books(context.Context) ([]BookResolver, error) {
	return []BookResolver{}, nil
}
