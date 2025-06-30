package handlers

import (
	"net/http"
	"dot-portal-go/internal/db/migrations"
)

// Migrate handles the migration request
func MigrateUp(w http.ResponseWriter, r *http.Request) {

	// Run the migration
	if err := migrations.MigrateUp(); err != nil {
		http.Error(w, "Migration failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Migration has been done successfully."}`))
}

// MigrateDown handles the rollback request
func MigrateDown(w http.ResponseWriter, r *http.Request) {

	// Run the migration down
	if err := migrations.MigrateDown(); err != nil {
		http.Error(w, "Rollback failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Migration rollback is successful."}`))
}