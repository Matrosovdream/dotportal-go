package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // or _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	// Replace with your DB config
	connStr := "postgres://postgres:password@db:5432/postgres?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database ping error:", err)
	}

	log.Println("Connected to the database")
}
