package resolver

import (
	"database/sql"
)

type Resolver struct {
	DB *sql.DB
}

func NewRoot(db *sql.DB) *Resolver {
	r := &Resolver{
		DB: db,
	}

	return r
}
