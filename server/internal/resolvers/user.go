package resolvers

import (
	"context"
	"database/sql"

	"github.com/bensaufley/aud-it/internal/entities/user"
)

type UserResolver struct {
	db *sql.DB
	user.User
}

func (u UserResolver) Lists(ctx context.Context) ([]ListResolver, error) {
	return []ListResolver{}, nil
}
