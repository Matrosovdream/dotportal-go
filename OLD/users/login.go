package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	//"fmt"
	"dot-portal-go/auth"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input LoginInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		var userID int
		var storedPassword string
		err := db.QueryRow("SELECT id, password FROM users WHERE email = $1", input.Email).
			Scan(&userID, &storedPassword)

		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		if input.Password != storedPassword {
			http.Error(w, "Wrong password", http.StatusUnauthorized)
			return
		}

		token, err := auth.GenerateToken(userID, input.Email)
		if err != nil {
			http.Error(w, "Token generation failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
	}
}
