package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// NewConnection returns a new database connection.
func NewConnection() *sqlx.DB {
	fmt.Println("Connecting to database...")
	db := sqlx.MustConnect("mysql", "root@tcp(localhost:3306)/gobatchservicedb")
	return db
}

// CloseConnection must be used to close the database connection
func CloseConnection(connection *sqlx.DB) error {
	fmt.Println("Closing connection ...")
	return connection.Close()
}
