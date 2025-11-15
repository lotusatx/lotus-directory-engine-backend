package main

import (
	"log"
	
	"github.com/lotusatx/lotus-directory-engine-backend/api"
	"github.com/lotusatx/lotus-directory-engine-backend/handlers"
	"github.com/lotusatx/lotus-directory-engine-backend/secrets"
)

func main() {
	LoadEnvFile()

	// Initialize secret manager (uses environment variables by default)
	secretManager := secrets.NewSecretManager()

	// Get database connection string
	cs, err := secretManager.GetConnectionString()
	if err != nil {
		log.Fatalf("Failed to get database connection string: %v", err)
	}

	// Initialize database connection and migration using existing db_handler
	db, err := handlers.ConfigureDbConnection(cs)
	if err != nil {
		log.Fatalf("Failed to configure database: %v", err)
	}

	// Get server port
	port := getEnvOrDefault("PORT", "8080")

	// Create and start the API server
	server := api.NewServer(db)
	
	log.Printf("Starting Lotus Directory Engine API server...")
	if err := server.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}