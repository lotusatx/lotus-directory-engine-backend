package api

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/lotusatx/lotus-directory-engine-backend/handlers"
	"github.com/lotusatx/lotus-directory-engine-backend/models"
	"gorm.io/gorm"
)

type GroupAPI struct {
	DB *gorm.DB
}

type AddUsersRequest struct {
	UserIDs []string `json:"user_ids"`
}

type AddUserRequest struct {
	UserID string `json:"user_id"`
}

// CreateGroup handles POST /api/groups
func (ga *GroupAPI) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group models.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.CreateGroup(ga.DB, &group); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
}

// GetGroup handles GET /api/groups/{id}
func (ga *GroupAPI) GetGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["id"]

	group, err := handlers.GetGroupByID(ga.DB, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

// GetAllGroups handles GET /api/groups
func (ga *GroupAPI) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := handlers.GetAllGroups(ga.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

// UpdateGroup handles PUT /api/groups/{id}
func (ga *GroupAPI) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["id"]

	var group models.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Ensure the ID matches the URL parameter
	group.ID = groupID

	if err := handlers.UpdateGroup(ga.DB, &group); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

// DeleteGroup handles DELETE /api/groups/{id}
func (ga *GroupAPI) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["id"]

	if err := handlers.DeleteGroup(ga.DB, groupID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AddUserToGroup handles POST /api/groups/{id}/users
func (ga *GroupAPI) AddUserToGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["id"]

	var req AddUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.AddUserToGroup(ga.DB, groupID, req.UserID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AddUsersToGroup handles POST /api/groups/{id}/users/bulk
func (ga *GroupAPI) AddUsersToGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["id"]

	var req AddUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.AddUsersToGroup(ga.DB, groupID, req.UserIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RemoveUserFromGroup handles DELETE /api/groups/{id}/users/{userId}
func (ga *GroupAPI) RemoveUserFromGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["id"]
	userID := vars["userId"]

	if err := handlers.RemoveUserFromGroup(ga.DB, groupID, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RemoveUsersFromGroup handles DELETE /api/groups/{id}/users/bulk
func (ga *GroupAPI) RemoveUsersFromGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["id"]

	var req AddUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.RemoveUsersFromGroup(ga.DB, groupID, req.UserIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetGroupMembers handles GET /api/groups/{id}/members
func (ga *GroupAPI) GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["id"]

	members, err := handlers.GetGroupMembers(ga.DB, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]string{"members": members})
}

// GetUserGroups handles GET /api/users/{userId}/groups
func (ga *GroupAPI) GetUserGroups(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	groups, err := handlers.GetUserGroups(ga.DB, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

// RegisterGroupRoutes registers all group-related routes
func (ga *GroupAPI) RegisterGroupRoutes(router *mux.Router) {
	groupRouter := router.PathPrefix("/groups").Subrouter()
	
	// Basic CRUD
	groupRouter.HandleFunc("", ga.CreateGroup).Methods("POST")
	groupRouter.HandleFunc("", ga.GetAllGroups).Methods("GET")
	groupRouter.HandleFunc("/{id}", ga.GetGroup).Methods("GET")
	groupRouter.HandleFunc("/{id}", ga.UpdateGroup).Methods("PUT")
	groupRouter.HandleFunc("/{id}", ga.DeleteGroup).Methods("DELETE")
	
	// Member management
	groupRouter.HandleFunc("/{id}/users", ga.AddUserToGroup).Methods("POST")
	groupRouter.HandleFunc("/{id}/users/bulk", ga.AddUsersToGroup).Methods("POST")
	groupRouter.HandleFunc("/{id}/users/{userId}", ga.RemoveUserFromGroup).Methods("DELETE")
	groupRouter.HandleFunc("/{id}/users/bulk", ga.RemoveUsersFromGroup).Methods("DELETE")
	groupRouter.HandleFunc("/{id}/members", ga.GetGroupMembers).Methods("GET")
	
	// User-group relationships
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/{userId}/groups", ga.GetUserGroups).Methods("GET")
}
