package database

import (
	"context"
	"database/sql"
)

type Connection interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func X() {
	ctx := context.Background()
	e, _ := sql.Open("postgres", "user=youruser dbname=yourdb sslmode=disable")
	// Example usage of context-aware methods
	e.ExecContext(ctx, "SELECT 1")
}
