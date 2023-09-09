package db

import (
	"context"
	"database/sql"
	"go-gin-sqlc-template/adapter/db/playground"
)

type TransactionFunc func(querier playground.Querier) error

type QueryWrapper interface {
	playground.Querier
	Transactionally(ctx context.Context, fn TransactionFunc) error
}

type queryWrapper struct {
	*sql.DB
	*playground.Queries
}

func (q *queryWrapper) Transactionally(ctx context.Context, fn TransactionFunc) error {
	tx, err := q.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qq := q.WithTx(tx)

	if err = fn(qq); err != nil {
		return err
	}

	return tx.Commit()
}

func WrapQuery(db *sql.DB, query *playground.Queries) QueryWrapper {
	return &queryWrapper{
		DB:      db,
		Queries: query,
	}
}
