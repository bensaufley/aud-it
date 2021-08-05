package resolver

import (
	"context"

	"github.com/bensaufley/aud-it/internal/resolvers"
	"github.com/graph-gophers/graphql-go"
)

func (r *Resolver) GetBook(ctx context.Context, args struct{ ID graphql.ID }) (*resolvers.NullableBookResolver, error) {
	return nil, nil
}

func (r *Resolver) GetBooks(ctx context.Context, args struct {
	Author *graphql.ID
	Genre  *graphql.ID
}) ([]resolvers.BookResolver, error) {
	return []resolvers.BookResolver{}, nil
}

func (r *Resolver) GetList(ctx context.Context, args struct{ ID graphql.ID }) (*resolvers.NullableListResolver, error) {
	return nil, nil
}

func (r *Resolver) GetLists(ctx context.Context) ([]resolvers.ListResolver, error) {
	return []resolvers.ListResolver{}, nil
}
