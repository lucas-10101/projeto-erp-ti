package database

import "database/sql"

type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	// Other common methods between sql.DB and sql.Tx
}
