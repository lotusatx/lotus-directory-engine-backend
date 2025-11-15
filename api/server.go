package api

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/gorm"
	"github.com/lotusatx/lotus-directory-engine-backend/secrets"
)

type Server struct {
	DB       *gorm.DB
	UserAPI  *UserAPI
	GroupAPI *GroupAPI
	RoleAPI  *RoleAPI
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		DB:       db,
		UserAPI:  &UserAPI{DB: db},
		GroupAPI: &GroupAPI{DB: db},
		RoleAPI:  &RoleAPI{DB: db},
	}
}

func (s *Server) SetupRoutes() http.Handler {
	router := mux.NewRouter()
	
	// API version prefix
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	
	// Register all routes
	s.UserAPI.RegisterUserRoutes(apiRouter)
	s.GroupAPI.RegisterGroupRoutes(apiRouter)
	s.RoleAPI.RegisterRoleRoutes(apiRouter)
	
	// Health check endpoint
	router.HandleFunc("/health", s.HealthCheck).Methods("GET")
	
	// Setup CORS
	corsOrigins := []string{"*"} // Default
	if origins := os.Getenv("CORS_ORIGINS"); origins != "" {
		corsOrigins = strings.Split(origins, ",")
	}
	
	c := cors.New(cors.Options{
		AllowedOrigins: corsOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})
	
	return c.Handler(router)
}

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok","service":"lotus-directory-engine"}`))
}

func (s *Server) loadTLSConfig() (*tls.Config, error) {
	// Check for PFX/P12 certificate first
	pfxFile := os.Getenv("TLS_PFX_FILE")
	
	if pfxFile != "" {
		// Get PFX password from secret manager
		secretManager := secrets.NewSecretManager()
		pfxPassword, err := secretManager.GetTLSPassword()
		if err != nil {
			log.Printf("Warning: Could not get PFX password: %v", err)
			pfxPassword = "" // Try without password
		}
		
		// Load PFX certificate
		cert, err := loadPFXCertificate(pfxFile, pfxPassword)
		if err != nil {
			return nil, err
		}
		
		return &tls.Config{
			Certificates: []tls.Certificate{cert},
		}, nil
	}
	
	// Fall back to separate cert and key files
	certFile := os.Getenv("TLS_CERT_FILE")
	keyFile := os.Getenv("TLS_KEY_FILE")
	
	if certFile != "" && keyFile != "" {
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, err
		}
		
		return &tls.Config{
			Certificates: []tls.Certificate{cert},
		}, nil
	}
	
	return nil, nil // No SSL configuration
}

func loadPFXCertificate(pfxFile, password string) (tls.Certificate, error) {
	// Note: Go's standard library doesn't have built-in PFX support
	// This is a placeholder - you would need to use a third-party library
	// like "software.sslmate.com/src/go-pkcs12" for PFX support
	
	log.Printf("PFX certificate support requires additional library")
	log.Printf("Consider converting PFX to PEM format:")
	log.Printf("  openssl pkcs12 -in %s -out cert.pem -clcerts -nokeys", pfxFile)
	log.Printf("  openssl pkcs12 -in %s -out key.pem -nocerts -nodes", pfxFile)
	
	return tls.Certificate{}, nil
}

func (s *Server) Start(port string) error {
	handler := s.SetupRoutes()
	
	// Try to load TLS configuration
	tlsConfig, err := s.loadTLSConfig()
	if err != nil {
		log.Printf("TLS configuration error: %v", err)
		return err
	}
	
	if tlsConfig != nil {
		// HTTPS mode
		server := &http.Server{
			Addr:      ":" + port,
			Handler:   handler,
			TLSConfig: tlsConfig,
		}
		
		log.Printf("Starting HTTPS server on port %s", port)
		log.Printf("Health check available at: https://localhost:%s/health", port)
		log.Printf("API endpoints available at: https://localhost:%s/api/v1/", port)
		
		return server.ListenAndServeTLS("", "")
	} else {
		// HTTP mode
		log.Printf("Starting HTTP server on port %s", port)
		log.Printf("Health check available at: http://localhost:%s/health", port)
		log.Printf("API endpoints available at: http://localhost:%s/api/v1/", port)
		log.Printf("Note: For production, configure TLS certificates for HTTPS")
		
		return http.ListenAndServe(":"+port, handler)
	}
}