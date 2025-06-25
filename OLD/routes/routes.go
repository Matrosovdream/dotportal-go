package routes

import (
	"github.com/gorilla/mux"
	"dot-portal-go/db"
	"dot-portal-go/auth"
	"dot-portal-go/users"
	"net/http"
)

func RegisterRoutes(router *mux.Router) {

	// Auth
	router.HandleFunc("/register", users.RegisterHandler(db.DB)).Methods("GET", "POST")
	router.HandleFunc("/login", users.LoginHandler(db.DB)).Methods("POST")

	// User routes
	protected := router.PathPrefix("/users").Subrouter()
	protected.Use(auth.JWTMiddleware) // JWT authentication middleware
	protected.HandleFunc("", users.GetUsersHandler(db.DB)).Methods("GET")
	protected.HandleFunc("/{id:[0-9]+}", users.DeleteUserHandler(db.DB)).Methods("GET", "DELETE")

}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Dot Portal API!"))
}

