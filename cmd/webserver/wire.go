//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-gin-sqlc-template/adapter/db"
	"go-gin-sqlc-template/adapter/db/playground"
	"go-gin-sqlc-template/adapter/router"
	"go-gin-sqlc-template/app/author"
	"go-gin-sqlc-template/domain"
	"net/http"

	"github.com/google/wire"
)

var (
	queryProvider = wire.NewSet(
		db.OpenMySQL,
		wire.Bind(new(playground.DBTX), new(*sql.DB)),
		playground.New,
	)

	authorUsecaseProvider = wire.NewSet(
		wire.Bind(new(domain.AuthorUsecase), new(*author.Usecase)),
		author.NewAuthorUsecase,
	)

	usecasesProvider = wire.NewSet(
		authorUsecaseProvider,
		wire.Struct(new(domain.Usecases), "*"),
	)
)

func serverProvider(handler http.Handler) *http.Server {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8000),
		Handler: handler,
	}

	return server
}

func initApp(ctx context.Context, dsn string) (*App, error) {
	wire.Build(
		queryProvider,
		usecasesProvider,
		router.InitHandler,
		serverProvider,
		wire.Struct(new(App), "*"),
	)

	return &App{}, nil
}
