package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	//"fmt"
)


func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var input RegisterInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate the input
		if err := validateRegisterInput(input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Check if the email already exists
		existUser, _ := findUserByEmail(db, input.Email)

		if( existUser != nil && existUser.Email != "" ) {
			http.Error(w, "Email already exists" + existUser.Email, http.StatusConflict)
			return
		}

		// Create the user
		errCreate := createUser(db, input.Email, input.Login, input.Password)

		if errCreate != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		writeHeaderContent(w, http.StatusCreated, "application/json", map[string]string{
			"message": "User created successfully",
			"email":   input.Email,
			"login":   input.Login,
		})
	}
}

func GetUsersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		users, err := getAllUsers(db)
		if err != nil {
			http.Error(w, "Error fetching users: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if len(users) == 0 {
			writeHeaderContent(w, http.StatusOK, "application/json", map[string]string{
				"message": "No users found",
			})
			return
		} else {
			writeHeaderContent(w, http.StatusOK, "application/json", users)
		}
	}
}

func DeleteUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		// Delete the user by ID
		errDelete := deleteUser(db, id)
		if errDelete != nil {
			http.Error(w, "Error deleting user: "+errDelete.Error(), http.StatusInternalServerError)
			return
		}

		writeHeaderContent(w, http.StatusOK, "application/json", map[string]string{
			"message": "User deleted successfully",
		})
	}
}

// Write writeHeaderContent in a json format
func writeHeaderContent(w http.ResponseWriter, status int, contentType string, data interface{}) {

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	if data != nil {
		response, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
		w.Write(response)
	}
	
}