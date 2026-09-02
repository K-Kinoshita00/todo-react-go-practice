package repository

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func getDSN() string {
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresDB := os.Getenv("POSTGRES_DB")
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", postgresUser, postgresPassword, postgresHost, postgresPort, postgresDB)
}

func dbSetup(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", getDSN())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	err = db.Ping()
	if err != nil {
		t.Fatalf("Ping: %v", err)
	}
	return db
}