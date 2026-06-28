package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// NewConnection returns a new database connection.
func NewConnection() *sqlx.DB {
	fmt.Println("Connecting to database...")
	db := sqlx.MustConnect("mysql", DSN())
	// Bound the pool so concurrent gobatch partitions queue for a connection
	// instead of exhausting MySQL's max_connections (Error 1040).
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	return db
}

// CloseConnection must be used to close the database connection
func CloseConnection(connection *sqlx.DB) error {
	fmt.Println("Closing connection ...")
	return connection.Close()
}
