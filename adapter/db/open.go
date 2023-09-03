package db

import (
	"context"
	"database/sql"
	"go-gin-sqlc-template/domain"

	_ "github.com/go-sql-driver/mysql"
)

// OpenMySQL connects to a MySQL database with the specified DSN (Data Source Name).
// This function will also ping the database to ensure it is accessible.
//
// Parameters:
// - ctx: Context to allow for potentially cancellable database operations. This does not cancel the connection itself.
// - dsn: A string containing the data source name. The DSN should be formatted using MySQL-specific syntax.
//
// Returns:
// On success, a pointer to a sql.DB (the core component of the package) and a nil error.
// On failure, returns a nil sql.DB and an error object indicating the reason for the failure.
//
// Example:
// db, err := OpenMySQL(context.Background(), "user:pass@/dbname")
//
//	if err != nil {
//	    log.Fatal(err)
//	}
func OpenMySQL(ctx context.Context, dbConfig domain.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConfig.DSN)
	if err != nil {
		return nil, err
	}
	// Use PingContext to check if the database is available and accessible.
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
