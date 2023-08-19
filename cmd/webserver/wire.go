package main

import (
	"context"
	"database/sql"
	"go-chi-sqlc-template/adapter/db"
	"go-chi-sqlc-template/adapter/db/playground"

	"github.com/google/wire"
)

var (
	queryProvider = wire.NewSet(
		db.OpenMySQL,
		wire.Bind(new(playground.DBTX), new(*sql.DB)),
		playground.New,
	)
)

func initApp(ctx context.Context, dsn string) (*App, error) {
	wire.Build(
		queryProvider,
		wire.Struct(new(App), "*"),
	)

	return &App{}, nil
}
