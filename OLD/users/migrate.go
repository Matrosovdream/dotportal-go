package users

import (
	"database/sql"
	"log"
)

func Migrate(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		login VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Could not run user migration: %v", err)
	}

	log.Println("Users table migrated")
}
