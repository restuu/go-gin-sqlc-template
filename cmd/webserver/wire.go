//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-gin-sqlc-template/adapter/config"
	"go-gin-sqlc-template/adapter/db"
	"go-gin-sqlc-template/adapter/db/playground"
	"go-gin-sqlc-template/adapter/router"
	"go-gin-sqlc-template/app/author"
	"go-gin-sqlc-template/domain"
	"net/http"

	"github.com/google/wire"
)

var (
	queryPlaygroundProvider = wire.NewSet(
		wire.FieldsOf(new(*domain.Config), "DB"),
		wire.FieldsOf(new(domain.Databases), "Playground"),
		db.OpenMySQL,
		wire.Bind(new(playground.DBTX), new(*sql.DB)),
		playground.New,
		db.WrapQuery,
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

func initApp(ctx context.Context) (*App, error) {
	wire.Build(
		config.ReadConfig,
		queryPlaygroundProvider,
		usecasesProvider,
		router.InitHandler,
		serverProvider,
		wire.Struct(new(App), "*"),
	)

	return &App{}, nil
}
