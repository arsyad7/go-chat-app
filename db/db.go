package db

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	conn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres", "postgres", "localhost", "5432", "go-chat-app",
	)

	// Open a connection to the database.
	db, err := sql.Open("postgres", conn)
	if err != nil {
		slog.Error(fmt.Sprintf("postgres: conn, err : %v", err.Error()))
	}

	if err = db.Ping(); err != nil {
		slog.Error(fmt.Sprintf("postgres: ping, err : %v", err.Error()))
	}
	fmt.Println("Connected to the database")

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
