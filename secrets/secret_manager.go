package secrets

import (
	"fmt"
	"os"
	"strings"
)

// SecretManager handles application secrets using environment variables
// For Kubernetes deployments, use Kubernetes secrets mounted as env vars or files
type SecretManager struct{}

// NewSecretManager creates a new secret manager
func NewSecretManager() *SecretManager {
	return &SecretManager{}
}

// GetSecret retrieves a secret from environment variables
func (sm *SecretManager) GetSecret(secretName string) (string, error) {
	value := os.Getenv(secretName)
	if value == "" {
		return "", fmt.Errorf("environment variable '%s' not set", secretName)
	}
	return value, nil
}

// GetConnectionString builds database connection string
func (sm *SecretManager) GetConnectionString() (string, error) {
	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		return "", fmt.Errorf("CONNECTION_STRING environment variable not set")
	}
	
	// If connection string already contains password, return as-is
	if !strings.Contains(connectionString, ":@") {
		return connectionString, nil
	}
	
	// Get password from environment variable
	passwordKey := getEnvOrDefault("DB_PASSWORD_KEY", "DB_PASSWORD")
	password, err := sm.GetSecret(passwordKey)
	if err != nil {
		return "", fmt.Errorf("failed to get database password: %w", err)
	}
	
	// Insert password into connection string
	connectionString = strings.Replace(connectionString, ":@", ":"+password+"@", 1)
	return connectionString, nil
}

// Helper functions for common secrets
func (sm *SecretManager) GetDatabasePassword() (string, error) {
	key := getEnvOrDefault("DB_PASSWORD_KEY", "DB_PASSWORD")
	return sm.GetSecret(key)
}

func (sm *SecretManager) GetJWTSecret() (string, error) {
	key := getEnvOrDefault("JWT_SECRET_KEY", "JWT_SECRET")
	return sm.GetSecret(key)
}

func (sm *SecretManager) GetTLSPassword() (string, error) {
	key := getEnvOrDefault("TLS_PASSWORD_KEY", "TLS_PASSWORD")
	return sm.GetSecret(key)
}

func (sm *SecretManager) GetAdminPassword() (string, error) {
	key := getEnvOrDefault("ADMIN_PASSWORD_KEY", "LDE_ADMIN_PASS")
	return sm.GetSecret(key)
}

// getEnvOrDefault returns environment variable value or default
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}