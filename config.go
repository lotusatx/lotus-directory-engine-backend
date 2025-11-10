package main

import (
	"fmt"
	"os"
	"strings"
)

// getEnvOrDefault returns the value of an environment variable or a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// LoadEnvFile loads environment variables from .env file if it exists
func LoadEnvFile() {
	const envFileName = ".env"

	// Check if .env file exists
	if _, err := os.Stat(envFileName); os.IsNotExist(err) {
		return // File doesn't exist, skip loading
	}

	// Read file contents
	content, err := os.ReadFile(envFileName)
	if err != nil {
		fmt.Printf("Warning: could not read %s: %v\n", envFileName, err)
		return
	}

	// Parse and set environment variables
	if err := parseEnvContent(string(content)); err != nil {
		fmt.Printf("Warning: error parsing %s: %v\n", envFileName, err)
	}
}

// parseEnvContent parses environment file content and sets variables
func parseEnvContent(content string) error {
	lines := strings.Split(content, "\n")
	for lineNum, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse key=value pairs
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid format on line %d: %q", lineNum+1, line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Only set if not already set in environment
		if os.Getenv(key) == "" {
			os.Setenv(key, value)
		}
	}
	return nil
}