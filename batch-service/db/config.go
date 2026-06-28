package db

import (
	"fmt"
	"os"
)

// getenv returns the value of the environment variable named by key,
// or fallback when the variable is unset or empty.
func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// DSN builds the MySQL data source name from environment variables, falling
// back to local development defaults. It is shared by the application
// connection and the gobatch metadata store, which target the same database.
func DSN() string {
	user := getenv("DB_USERNAME", "root")
	password := getenv("DB_PASSWORD", "")
	host := getenv("DB_HOST", "127.0.0.1")
	port := getenv("DB_PORT", "3306")
	name := getenv("DB_NAME", "gobatchservicedb")

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		user, password, host, port, name,
	)
}
