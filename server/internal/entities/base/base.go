package base

import (
	"github.com/bensaufley/aud-it/internal/ulid"
	"github.com/graph-gophers/graphql-go"
)

type Entity struct {
	ULID string
}

func (b *Entity) ID() graphql.ID {
	return graphql.ID(b.ULID)
}

func (b *Entity) GetULID() string {
	if b.ULID != "" {
		return b.ULID
	}
	b.ULID = ulid.NewGenerator().New().String()
	return b.ULID
}

func StrPointer(str string) *string {
	return &str
}
