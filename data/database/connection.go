package database

import (
	"context"
	"database/sql"
	"erp/organization-api/utils"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	connection        *sql.DB
	mongodbConnection *mongo.Client
)

// SQL Default connection

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
	connection, err = sql.Open(utils.ApplicationProperties.DatabaseDriver, utils.ApplicationProperties.DatabaseConnectionString)
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

// NoSQL Default connection

func CreateMongoDBConnection() *mongo.Client {

	options := options.Client().
		ApplyURI(utils.ApplicationProperties.MongoDBConnectionString)
		// SetAppName(utils.ApplicationProperties.ApplicationName).
		// SetMaxPoolSize(100).
		// SetMinPoolSize(2).
		// SetWriteConcern(writeconcern.Majority()).
		// SetTimeout(time.Second * 10).
		// SetServerSelectionTimeout(time.Second * 10)

	var err error
	mongodbConnection, err = mongo.Connect(options)
	if err != nil {
		panic(err)
	}

	if err = mongodbConnection.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}

	return mongodbConnection
}

func GetMongoDBConnection() *mongo.Client {
	return mongodbConnection
}
