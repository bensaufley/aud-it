package resolvers

import (
	"context"
	"database/sql"

	"github.com/bensaufley/aud-it/internal/entities/list"
)

type ListResolver struct {
	db *sql.DB
	list.List
}

type NullableListResolver ListResolver

func (l *NullableListResolver) User(ctx context.Context) (UserResolver, error) {
	if l == nil {
		return UserResolver{}, nil
	}
	return ListResolver(*l).User(ctx)
}

func (l ListResolver) User(ctx context.Context) (UserResolver, error) {
	return UserResolver{db: l.db}, nil
}

func (l *NullableListResolver) Books(ctx context.Context) ([]BookResolver, error) {
	if l == nil {
		return []BookResolver{}, nil
	}
	return ListResolver(*l).Books(ctx)
}

func (l ListResolver) Books(ctx context.Context) ([]BookResolver, error) {
	return []BookResolver{}, nil
}
