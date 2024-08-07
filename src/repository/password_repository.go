package repository

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

type PasswordRepository struct {
    DB *sql.DB
}

func NewPasswordRepository(db *sql.DB) *PasswordRepository {
    return &PasswordRepository{DB: db}
}

func (r *PasswordRepository) LogRequestResponse(request, response string) {
    query := `INSERT INTO password_logs (request, response) VALUES ($1, $2)`
    if _, err := r.DB.Exec(query, request, response); err != nil {
        log.Printf("Failed to log request and response: %v", err)
    }
}
