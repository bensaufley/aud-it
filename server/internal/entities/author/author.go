package author

import (
	"context"
	"database/sql"
	"strings"

	"github.com/bensaufley/aud-it/internal/entities/base"
)

type Author struct {
	*base.Entity
	FirstName *string
	LastName  string
}

// Order of names: lastName, firstName
func New(names ...string) *Author {
	lastName := ""
	var firstName *string
	if len(names) > 0 {
		lastName = names[0]
		if len(names) > 1 {
			firstName = &names[1]
		}
	}
	return &Author{
		Entity:    &base.Entity{},
		FirstName: firstName,
		LastName:  lastName,
	}
}

func Create(ctx context.Context, tx *sql.Tx, authors ...*Author) error {
	template := strings.Builder{}
	args := []interface{}{}
	for i, author := range authors {
		if i != 0 {
			template.WriteString(", ")
		}
		template.WriteString("(?, ?, ?)")
		args = append(args, author.GetULID(), author.FirstName, author.LastName)
	}
	if _, err := tx.ExecContext(
		ctx,
		`INSERT INTO authors (ulid, first_name, last_name) VALUES `+template.String(),
		args...,
	); err != nil {
		return err
	}
	return nil
}

func Read(ctx context.Context, db *sql.DB, ids ...string) ([]*Author, error) {
	template := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		template = append(template, "?")
		args[i] = id
	}
	rows, err := db.QueryContext(ctx, `SELECT (ulid, first_name, last_name) FROM authors WHERE ulid IN (`+strings.Join(template, ", ")+`)`, args...)
	authors := []*Author{}
	if err != nil {
		return []*Author{}, err
	}
	for rows.Next() {
		author := Author{Entity: &base.Entity{}}
		if err := rows.Scan(&author.ULID, &author.FirstName, &author.LastName); err != nil {
			return []*Author{}, err
		}
	}
	return authors, nil
}

func Update(ctx context.Context, tx *sql.Tx, authors ...*Author) error {
	for _, author := range authors {
		if _, err := tx.ExecContext(ctx, `UPDATE authors SET first_name=?, last_name=? WHERE ulid=?`, author.FirstName, author.LastName, author.ULID); err != nil {
			return err
		}
	}
	return nil
}

func Destroy(ctx context.Context, tx *sql.Tx, ids ...string) error {
	template := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		template[i] = "?"
		args[i] = id
	}
	_, err := tx.ExecContext(ctx, `DELETE FROM authors WHERE ulid IN (`+strings.Join(template, ", ")+`)`, args...)
	return err
}
