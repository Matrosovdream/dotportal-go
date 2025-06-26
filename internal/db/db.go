package db

import (
  "database/sql"
  "fmt"
  "os"

  _ "github.com/lib/pq"
)

var CONN *sql.DB

// Init initializes the database connection pool.
func Init() error {
	
  dsn := os.Getenv("DB_DSN")
  if dsn == "" {
    return fmt.Errorf("DB_DSN env is required")
  }
  var err error
  CONN, err = sql.Open("postgres", dsn)
  if err != nil {
    return err
  }
  return CONN.Ping()

}
