package storage

import (
    "database/sql"
    _ "github.com/lib/pq"
)

// ConectarDB establece conexión con PostgreSQL
func ConectarDB() (*sql.DB, error) {
    connStr := "user=postgres password=1234 dbname=libros sslmode=disable"
    return sql.Open("postgres", connStr)
}
