package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"

	"github.com/bensaufley/graphql-preact-starter/internal/ulid"
)

type TodoInput struct {
	Contents string
	Status   *string
}

func TodoStatusFromString(str *string) (TodoStatus, error) {
	if str == nil {
		return Unstarted, nil
	}

	switch *str {
	case "UNSTARTED":
		return Unstarted, nil
	case "IN_PROGRESS":
		return InProgress, nil
	case "ABANDONED":
		return Abandoned, nil
	case "COMPLETE":
		return Complete, nil
	case "DELETED":
		return Deleted, nil
	default:
		return Unstarted, errors.New("unrecognized TodoStatus")
	}
}

func (r *Resolver) AddTodo(ctx context.Context, args *struct{ Input TodoInput }) (*Todo, error) {
	status, err := TodoStatusFromString(args.Input.Status)
	if err != nil {
		return nil, err
	}
	ulid := ulid.NewGenerator().String()
	stmt, err := r.DB.PrepareContext(
		ctx,
		`INSERT INTO todos (ulid, contents, todo_status_id) VALUES (?, ?, (SELECT id FROM todo_statuses WHERE val = ?));`,
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not prepare statement")
	}
	if _, err = stmt.ExecContext(ctx, ulid, args.Input.Contents, status); err != nil {
		return nil, errors.Wrap(err, "could not execute insert")
	}

	todo := Todo{
		ID:       graphql.ID(ulid),
		Contents: args.Input.Contents,
		status:   status,
	}

	go func(r *Resolver, t Todo) {
		r.Subscriptions.todoAdded <- t
	}(r, todo)

	return &todo, nil
}

func (r *Resolver) DeleteTodo(ctx context.Context, args *struct{ ID graphql.ID }) (bool, error) {
	_, err := r.DB.ExecContext(ctx, `UPDATE todos SET todo_status_id=(SELECT id FROM todo_statuses WHERE val = 'Deleted') WHERE ulid = ?;`, args.ID)
	if err != nil {
		return false, errors.Wrap(err, "could not execute soft-delete")
	}

	go func(r *Resolver, id graphql.ID) {
		r.Subscriptions.todoDeleted <- id
	}(r, args.ID)

	return true, nil
}
