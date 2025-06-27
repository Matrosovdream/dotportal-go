package routes

import (
  "github.com/gorilla/mux"
  "dot-portal-go/internal/handlers"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	// Index page
	r.HandleFunc("/", handlers.Index).Methods("GET")

  // Registration
  r.HandleFunc("/register", handlers.Register).Methods("POST")

  // Login
  r.HandleFunc("/login", handlers.Login).Methods("POST")

	return r

}

