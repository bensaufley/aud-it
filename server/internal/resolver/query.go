package resolver

import (
	"context"
	"strings"

	"github.com/graph-gophers/graphql-go"
)

type TodoStatus string

const (
	Unstarted  TodoStatus = "Unstarted"
	InProgress TodoStatus = "In Progress"
	Abandoned  TodoStatus = "Abandoned"
	Complete   TodoStatus = "Complete"
	Deleted    TodoStatus = "Deleted"
)

func (t TodoStatus) ToGQLEnum() string {
	switch t {
	case InProgress:
		return "IN_PROGRESS"
	default:
		return strings.ToUpper(string(t))
	}
}

type Todo struct {
	ID       graphql.ID
	Contents string
	status   TodoStatus
}

func (r *Resolver) GetTodos(ctx context.Context) ([]Todo, error) {
	rows, err := r.DB.QueryContext(ctx, `SELECT t.ulid, contents, val FROM todos t JOIN todo_statuses s ON todo_status_id = s.id WHERE s.val != 'Deleted';`)
	response := []Todo{}
	if err != nil {
		return response, err
	}
	defer rows.Close()
	for rows.Next() {
		todo := Todo{}
		if err := rows.Scan(&todo.ID, &todo.Contents, &todo.status); err != nil {
			return []Todo{}, err
		}
		response = append(response, todo)
	}
	return response, nil
}

func (r *Resolver) GetTodo(ctx context.Context, args *struct{ ID string }) (*Todo, error) {
	rows := r.DB.QueryRowContext(ctx, `SELECT t.ulid, contents, val FROM todos t JOIN todo_statuses s ON todo_status_id = s.id WHERE t.ulid = ?`, args.ID)
	if err := rows.Err(); err != nil {
		return nil, err
	}
	todo := Todo{}
	if err := rows.Scan(&todo.ID, &todo.Contents, &todo.status); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t Todo) Status() string {
	return t.status.ToGQLEnum()
}
