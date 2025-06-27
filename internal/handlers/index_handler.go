package handlers

import (
	"net/http"	
	//"github.com/gorilla/mux"
)

// IndexHandler handles the root endpoint.
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Welcome to the Dot Portal API!"}`))
}