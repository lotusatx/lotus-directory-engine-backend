package handlers

import (
	"fmt"
	"strings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/lotusatx/lotus-directory-engine-backend/models"
)

func ConfigureDbConnection(connectionString string) (*gorm.DB, error) {
	// Ensure SSL is disabled for local/development PostgreSQL servers
	if !strings.Contains(connectionString, "sslmode=") {
		if strings.Contains(connectionString, "?") {
			connectionString += "&sslmode=disable"
		} else {
			connectionString += "?sslmode=disable"
		}
	}
	
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
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
	
	// Auto-migrate the schema
	err = migrateDatabase(db)
	if err != nil {
		fmt.Println("Database migration failed:", err)
		return nil, err
	}
	fmt.Println("Database migration completed successfully")
	
	return db, nil
}

func testDbConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func migrateDatabase(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Group{}, &models.Role{})
}

func Create