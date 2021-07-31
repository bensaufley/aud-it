package resolver

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"

	"github.com/bensaufley/graphql-preact-starter/internal/ulid"
)

type Subscriptions struct {
	todoAdded   chan Todo
	todoDeleted chan graphql.ID

	addedSubscribers   chan *addedSubscriber
	deletedSubscribers chan *deletedSubscriber
}

type addedSubscriber struct {
	stop   <-chan struct{}
	events chan<- Todo
}

type deletedSubscriber struct {
	stop   <-chan struct{}
	events chan<- graphql.ID
}

func (r *Resolver) TodoAdded(ctx context.Context) (<-chan Todo, error) {
	ch := make(chan Todo)
	r.Subscriptions.addedSubscribers <- &addedSubscriber{events: ch, stop: ctx.Done()}
	return ch, nil
}

func (r *Resolver) TodoDeleted(ctx context.Context) (<-chan graphql.ID, error) {
	ch := make(chan graphql.ID)
	r.Subscriptions.deletedSubscribers <- &deletedSubscriber{events: ch, stop: ctx.Done()}
	return ch, nil
}

func (r *Resolver) broadcastTodoChanges() {
	addedSubscribers := map[string]*addedSubscriber{}
	deletedSubscribers := map[string]*deletedSubscriber{}
	unsubscribe := make(chan string)
	ug := ulid.NewGenerator()

	for {
		select {
		case id := <-unsubscribe:
			delete(addedSubscribers, id)
		case s := <-r.Subscriptions.addedSubscribers:
			addedSubscribers[ug.String()] = s
		case s := <-r.Subscriptions.deletedSubscribers:
			deletedSubscribers[ug.String()] = s
		case e := <-r.Subscriptions.todoAdded:
			for id, s := range addedSubscribers {
				go func(id string, s *addedSubscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- e:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		case e := <-r.Subscriptions.todoDeleted:
			for id, s := range deletedSubscribers {
				go func(id string, s *deletedSubscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- e:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		}
	}
}
