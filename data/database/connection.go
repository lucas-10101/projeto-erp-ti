package database

import (
	"context"
	"database/sql"
)

var (
	connection *sql.DB
)

type Connection interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

func CreateConnection() {
	if connection != nil {
		panic("connection already exists")
	}

	var err error
	connection, err = sql.Open("postgres", "user=youruser dbname=yourdb sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = connection.Ping(); err != nil {
		panic(err)
	}
}

func GetUnderlingConnection() *sql.DB {
	return connection
}

func GetTransaction(ctx context.Context) (*sql.Tx, error) {
	return connection.BeginTx(ctx, nil)
}

func GetConnection() Connection {
	return connection
}
