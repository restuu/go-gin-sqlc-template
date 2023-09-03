package domain

import (
	"context"
	"go-gin-sqlc-template/adapter/db/playground"
)

// AuthorUsecase represents the author usecase contract
type AuthorUsecase interface {
	FindAllAuthors(ctx context.Context) ([]playground.Author, error)
}
