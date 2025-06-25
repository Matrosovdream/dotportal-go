package users

import (
	"database/sql"
	"fmt"
)

// Migrate creates the users table if it does not exist
func findUserByEmail(db *sql.DB, email string) (*User, error) {
	var user User

	query := "SELECT id, email, login, password FROM users WHERE email = $1"
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Login, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err // Other error
	}
	return &user, nil
}

func createUser(db *sql.DB, email, login, password string) error {
	_, err := db.Exec(
		"INSERT INTO users (email, login, password) VALUES ($1, $2, $3)",
		email, login, password,
	)
	return err
}

func validateRegisterInput(input RegisterInput) error {

	if input.Email == "" || input.Login == "" || input.Password == "" {
		return fmt.Errorf("missing fields: email, login, and password are required")
	}

	return nil

}

func getAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, email, login, password FROM users")
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Login, &user.Password); err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users: %w", err)
	}

	return users, nil
}

func getUserByEmail(db *sql.DB, email string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, email, login, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("error fetching user by email: %w", err)
	}
	return &user, nil
}

func deleteUser(db *sql.DB, userId int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userId)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}