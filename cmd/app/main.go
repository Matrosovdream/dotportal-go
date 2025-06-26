package main

import (
  "log"
  "net/http"

  "dot-portal-go/internal/db"
  "dot-portal-go/internal/routes"
)

func main() {

  // Init connection
  if err := db.Init(); err != nil {
    log.Fatal("DB init failed:", err)
  }
  defer db.CONN.Close()

  // Init router
  router := routes.NewRouter()

  // Init the server
  log.Println("Listening on :8080")
  log.Fatal(http.ListenAndServe(":8080", router))
}
