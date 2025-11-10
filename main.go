package main

import (
	"fmt"
	"os"
	//"github.com/lotusatx/lotus-directory-engine-backend/models"
	
	"github.com/lotusatx/lotus-directory-engine-backend/handlers"
)

func main() {
	LoadEnvFile()

	cs := os.Getenv("CONNECTION_STRING")

	db, err := handlers.ConfigureDbConnection(cs)
	if err != nil {
		fmt.Println("Error configuring database connection:", err)
		return
	}
	defer db.Close()

	fmt.Println("Database connection established successfully")
}