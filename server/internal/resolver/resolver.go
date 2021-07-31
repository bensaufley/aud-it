package resolver

import (
	"database/sql"

	"github.com/graph-gophers/graphql-go"
)

type Resolver struct {
	DB            *sql.DB
	Subscriptions Subscriptions
}

func NewRoot(db *sql.DB) *Resolver {
	r := &Resolver{
		DB: db,
		Subscriptions: Subscriptions{
			todoAdded:   make(chan Todo),
			todoDeleted: make(chan graphql.ID),

			addedSubscribers:   make(chan *addedSubscriber),
			deletedSubscribers: make(chan *deletedSubscriber),
		},
	}

	go r.broadcastTodoChanges()

	return r
}
