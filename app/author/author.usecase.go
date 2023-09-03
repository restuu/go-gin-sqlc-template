package author

import (
	"context"
	"go-gin-sqlc-template/adapter/db/playground"
	"sync"
)

// Usecase ...
type Usecase struct {
	queries *playground.Queries
}

var (
	authorUsecase     *Usecase
	authorUsecaseOnce sync.Once
)

// NewAuthorUsecase ...
func NewAuthorUsecase(queries *playground.Queries) *Usecase {
	authorUsecaseOnce.Do(func() {
		authorUsecase = &Usecase{
			queries: queries,
		}
	})

	return authorUsecase
}

func (a *Usecase) FindAllAuthors(ctx context.Context) ([]playground.Author, error) {
	return a.queries.FindAllAuthors(ctx)
}
