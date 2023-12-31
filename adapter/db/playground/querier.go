// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package playground

import (
	"context"
)

type Querier interface {
	FindAllAuthors(ctx context.Context) ([]Author, error)
	FindAllBooks(ctx context.Context) ([]Book, error)
	FindAuthorByID(ctx context.Context, id int64) (Author, error)
	FindAuthorsWithName(ctx context.Context, concat interface{}) ([]Author, error)
	FindBooksByAuthor(ctx context.Context, authorID int64) ([]Book, error)
	InsertAuthors(ctx context.Context, name string) (int64, error)
	InsertBook(ctx context.Context, arg InsertBookParams) (int64, error)
}

var _ Querier = (*Queries)(nil)
