package api

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/lotusatx/lotus-directory-engine-backend/handlers"
	"github.com/lotusatx/lotus-directory-engine-backend/models"
	"gorm.io/gorm"
)

type UserAPI struct {
	DB *gorm.DB
}

// CreateUser handles POST /api/users
func (ua *UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.CreateUser(ua.DB, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser handles GET /api/users/{id}
func (ua *UserAPI) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := handlers.GetUserByID(ua.DB, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetAllUsers handles GET /api/users
func (ua *UserAPI) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handlers.GetAllUsers(ua.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UpdateUser handles PUT /api/users/{id}
func (ua *UserAPI) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Ensure the ID matches the URL parameter
	user.ID = userID

	if err := handlers.UpdateUser(ua.DB, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUser handles DELETE /api/users/{id}
func (ua *UserAPI) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	if err := handlers.DeleteUser(ua.DB, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RegisterUserRoutes registers all user-related routes
func (ua *UserAPI) RegisterUserRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/users").Subrouter()
	
	userRouter.HandleFunc("", ua.CreateUser).Methods("POST")
	userRouter.HandleFunc("", ua.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", ua.GetUser).Methods("GET")
	userRouter.HandleFunc("/{id}", ua.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/{id}", ua.DeleteUser).Methods("DELETE")
}
