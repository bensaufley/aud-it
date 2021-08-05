package book

import (
	"context"
	"database/sql"
	"strings"

	"github.com/bensaufley/aud-it/internal/entities/author"
	"github.com/bensaufley/aud-it/internal/entities/base"
	"github.com/bensaufley/aud-it/internal/ulid"
)

type ImportStatus string

const (
	StatusPending  ImportStatus = "Pending"
	StatusComplete ImportStatus = "Complete"
)

type Book struct {
	*base.Entity
	Title        string
	Duration     *int32
	ImportStatus ImportStatus
}

type BookAuthor struct {
	*base.Entity
	Author *author.Author
	Meta   *string
}

func Create(ctx context.Context, tx *sql.Tx, book *Book, authors []*BookAuthor, genres []string) error {
	if _, err := tx.ExecContext(ctx, `INSERT INTO books (id, title, duration, import_status) VALUES (?, ?, ?, ?);`, book.GetULID(), book.Title, book.Duration, book.ImportStatus); err != nil {
		return err
	}
	gen := ulid.NewGenerator()
	bookAuthorTemplate := strings.Builder{}
	bookAuthorArgs := []interface{}{}
	for i, author := range authors {
		if i != 0 {
			bookAuthorTemplate.WriteString(", ")
		}
		bookAuthorTemplate.WriteString("(?, ?, ?)")
		bookAuthorArgs = append(bookAuthorArgs, gen.New().String(), book.ULID, author.Author.GetULID(), author.Meta)
	}
	if _, err := tx.ExecContext(
		ctx,
		`INSERT INTO book_authors (id, book_ulid, author_ulid, meta) VALUES `+bookAuthorTemplate.String(),
		bookAuthorArgs...,
	); err != nil {
		return err
	}
	return nil
}
