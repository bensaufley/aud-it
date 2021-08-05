package resolvers

import (
	"context"
	"database/sql"

	"github.com/bensaufley/aud-it/internal/entities/base"
	"github.com/bensaufley/aud-it/internal/entities/book"
)

type BookResolver struct {
	db *sql.DB
	book.Book
}

type NullableBookResolver BookResolver

type BookAuthorResolver struct {
	db *sql.DB
	*base.Entity
	Meta   *string
	author AuthorResolver
}

func (b *NullableBookResolver) Authors(ctx context.Context) ([]BookAuthorResolver, error) {
	if b == nil {
		return nil, nil
	}
	return BookResolver(*b).Authors(ctx)
}

func (b BookResolver) Authors(ctx context.Context) ([]BookAuthorResolver, error) {
	return []BookAuthorResolver{}, nil
}

func (b *NullableBookResolver) Genres(ctx context.Context) ([]GenreResolver, error) {
	if b == nil {
		return nil, nil
	}
	return BookResolver(*b).Genres(ctx)
}

func (b BookResolver) Genres(ctx context.Context) ([]GenreResolver, error) {
	return []GenreResolver{}, nil
}

func (ba BookAuthorResolver) Author(ctx context.Context) (AuthorResolver, error) {
	return AuthorResolver{db: ba.db}, nil
}

func (ba BookAuthorResolver) Book(ctx context.Context) (BookResolver, error) {
	return BookResolver{db: ba.db}, nil
}
