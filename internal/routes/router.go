package routes

import (
  //"net/http"
  //"fmt"

  "github.com/gorilla/mux"
  //"dot-portal-go/internal/db"
  "dot-portal-go/internal/handlers"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	// 1. Main page
	r.HandleFunc('/', register).
		Methods("GET").
		Headers("Content-Type", "application/json")	

	return r

}

/*
func NewRouter(userHandler *handlers.UserHandler) *mux.Router {
  r := mux.NewRouter()

  // 2. Register endpoint
  r.HandleFunc("/register", userHandler.Register).
    Methods("POST").
    Headers("Content-Type", "application/json")

  // 3. Login endpoint
  r.HandleFunc("/login", userHandler.Login).
    Methods("POST").
    Headers("Content-Type", "application/json")

  return r
}
*/
