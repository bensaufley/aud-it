package resolvers

import (
	"context"
	"database/sql"

	"github.com/bensaufley/aud-it/internal/entities/author"
)

type AuthorResolver struct {
	db *sql.DB
	author.Author
}

func (a AuthorResolver) Books(ctx context.Context) ([]BookAuthorResolver, error) {
	return []BookAuthorResolver{}, nil
}
