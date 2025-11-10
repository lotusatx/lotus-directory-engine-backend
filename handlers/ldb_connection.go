package handlers

import (
	"database/sql"
	"fmt"
	"strings"
	_ "github.com/lib/pq"
)

func ConfigureDbConnection(connectionString string) (*sql.DB, error) {
	// Ensure SSL is disabled for local/development PostgreSQL servers
	if !strings.Contains(connectionString, "sslmode=") {
		if strings.Contains(connectionString, "?") {
			connectionString += "&sslmode=disable"
		} else {
			connectionString += "?sslmode=disable"
		}
	}
	
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Failed to open database connection:", err)
		return nil, err
	}
	err = testDbConnection(db)
	if err != nil {
		fmt.Println("Database connection test failed:", err)
		return nil, err
	}
	fmt.Println("Database connection configured successfully")
	return db, nil
}

func testDbConnection(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}