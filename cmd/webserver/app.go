package main

import (
	"context"
	"database/sql"
	"go-gin-sqlc-template/adapter/db/playground"
	"log"
	"net/http"
)

type App struct {
	db      *sql.DB
	queries *playground.Queries
	server  *http.Server
}

func (a *App) start() {
	err := a.server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) shutdown(ctx context.Context) error {
	log.Println("closing database connection...")
	err := a.db.Close()
	if err != nil {
		return err
	}

	log.Println("shutting down server...")
	return a.server.Shutdown(ctx)
}
