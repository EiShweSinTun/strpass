package model

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func SetupDB() (*sql.DB, error) {
    // Use the environment variables if defined or hardcode if running locally
    connStr := "user=user1 password=rootpassword dbname=strpass host=postgres port=5432 sslmode=disable"

    postgres, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    // Test the connection
    if err := postgres.Ping(); err != nil {
        return nil, err
    }

    // Create the table if it does not exist
    if _, err := postgres.Exec(`CREATE TABLE IF NOT EXISTS password_logs (
        id SERIAL PRIMARY KEY,
        request TEXT,
        response TEXT
    )`); err != nil {
        return nil, err
    }

    return postgres, nil
}
