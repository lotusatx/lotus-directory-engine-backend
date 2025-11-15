package api

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/lotusatx/lotus-directory-engine-backend/handlers"
	"github.com/lotusatx/lotus-directory-engine-backend/models"
	"gorm.io/gorm"
)

type RoleAPI struct {
	DB *gorm.DB
}

type AddGroupsRequest struct {
	GroupIDs []string `json:"group_ids"`
}

type AddGroupRequest struct {
	GroupID string `json:"group_id"`
}

type AssignRoleRequest struct {
	RoleID string `json:"role_id"`
}

type AssignRolesRequest struct {
	RoleIDs []string `json:"role_ids"`
}

type BulkAssignRequest struct {
	UserIDs []string `json:"user_ids"`
}

// CreateRole handles POST /api/roles
func (ra *RoleAPI) CreateRole(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.CreateRole(ra.DB, &role); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(role)
}

// GetRole handles GET /api/roles/{id}
func (ra *RoleAPI) GetRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID := vars["id"]

	role, err := handlers.GetRoleByID(ra.DB, roleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

// GetAllRoles handles GET /api/roles
func (ra *RoleAPI) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := handlers.GetAllRoles(ra.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}

// UpdateRole handles PUT /api/roles/{id}
func (ra *RoleAPI) UpdateRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID := vars["id"]

	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Ensure the ID matches the URL parameter
	role.ID = roleID

	if err := handlers.UpdateRole(ra.DB, &role); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

// DeleteRole handles DELETE /api/roles/{id}
func (ra *RoleAPI) DeleteRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID := vars["id"]

	if err := handlers.DeleteRole(ra.DB, roleID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AddGroupToRole handles POST /api/roles/{id}/groups
func (ra *RoleAPI) AddGroupToRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID := vars["id"]

	var req AddGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.AddGroupToRole(ra.DB, roleID, req.GroupID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AddGroupsToRole handles POST /api/roles/{id}/groups/bulk
func (ra *RoleAPI) AddGroupsToRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID := vars["id"]

	var req AddGroupsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.AddGroupsToRole(ra.DB, roleID, req.GroupIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RemoveGroupFromRole handles DELETE /api/roles/{id}/groups/{groupId}
func (ra *RoleAPI) RemoveGroupFromRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID := vars["id"]
	groupID := vars["groupId"]

	if err := handlers.RemoveGroupFromRole(ra.DB, roleID, groupID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RemoveGroupsFromRole handles DELETE /api/roles/{id}/groups/bulk
func (ra *RoleAPI) RemoveGroupsFromRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID := vars["id"]

	var req AddGroupsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.RemoveGroupsFromRole(ra.DB, roleID, req.GroupIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AssignRoleToUser handles POST /api/users/{userId}/roles
func (ra *RoleAPI) AssignRoleToUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	var req AssignRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.AssignRoleToUser(ra.DB, userID, req.RoleID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AssignRolesToUser handles POST /api/users/{userId}/roles/bulk
func (ra *RoleAPI) AssignRolesToUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	var req AssignRolesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.AssignRolesToUser(ra.DB, userID, req.RoleIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RemoveRoleFromUser handles DELETE /api/users/{userId}/roles/{roleId}
func (ra *RoleAPI) RemoveRoleFromUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]
	roleID := vars["roleId"]

	if err := handlers.RemoveRoleFromUser(ra.DB, userID, roleID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RemoveRolesFromUser handles DELETE /api/users/{userId}/roles/bulk
func (ra *RoleAPI) RemoveRolesFromUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	var req AssignRolesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.RemoveRolesFromUser(ra.DB, userID, req.RoleIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// BulkAssignRoleToUsers handles POST /api/roles/{id}/users/bulk
func (ra *RoleAPI) BulkAssignRoleToUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID := vars["id"]

	var req BulkAssignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.BulkAssignRoleToUsers(ra.DB, roleID, req.UserIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// BulkRemoveRoleFromUsers handles DELETE /api/roles/{id}/users/bulk
func (ra *RoleAPI) BulkRemoveRoleFromUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID := vars["id"]

	var req BulkAssignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := handlers.BulkRemoveRoleFromUsers(ra.DB, roleID, req.UserIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetUserRoles handles GET /api/users/{userId}/roles
func (ra *RoleAPI) GetUserRoles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	roles, err := handlers.GetUserRoles(ra.DB, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}

// GetRoleGroups handles GET /api/roles/{id}/groups
func (ra *RoleAPI) GetRoleGroups(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID := vars["id"]

	groups, err := handlers.GetRoleGroups(ra.DB, roleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]string{"groups": groups})
}

// RegisterRoleRoutes registers all role-related routes
func (ra *RoleAPI) RegisterRoleRoutes(router *mux.Router) {
	roleRouter := router.PathPrefix("/roles").Subrouter()
	
	// Basic CRUD
	roleRouter.HandleFunc("", ra.CreateRole).Methods("POST")
	roleRouter.HandleFunc("", ra.GetAllRoles).Methods("GET")
	roleRouter.HandleFunc("/{id}", ra.GetRole).Methods("GET")
	roleRouter.HandleFunc("/{id}", ra.UpdateRole).Methods("PUT")
	roleRouter.HandleFunc("/{id}", ra.DeleteRole).Methods("DELETE")
	
	// Group management for roles
	roleRouter.HandleFunc("/{id}/groups", ra.AddGroupToRole).Methods("POST")
	roleRouter.HandleFunc("/{id}/groups/bulk", ra.AddGroupsToRole).Methods("POST")
	roleRouter.HandleFunc("/{id}/groups/{groupId}", ra.RemoveGroupFromRole).Methods("DELETE")
	roleRouter.HandleFunc("/{id}/groups/bulk", ra.RemoveGroupsFromRole).Methods("DELETE")
	roleRouter.HandleFunc("/{id}/groups", ra.GetRoleGroups).Methods("GET")
	
	// Bulk user assignment for roles
	roleRouter.HandleFunc("/{id}/users/bulk", ra.BulkAssignRoleToUsers).Methods("POST")
	roleRouter.HandleFunc("/{id}/users/bulk", ra.BulkRemoveRoleFromUsers).Methods("DELETE")
	
	// User-role relationships
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/{userId}/roles", ra.AssignRoleToUser).Methods("POST")
	userRouter.HandleFunc("/{userId}/roles/bulk", ra.AssignRolesToUser).Methods("POST")
	userRouter.HandleFunc("/{userId}/roles/{roleId}", ra.RemoveRoleFromUser).Methods("DELETE")
	userRouter.HandleFunc("/{userId}/roles/bulk", ra.RemoveRolesFromUser).Methods("DELETE")
	userRouter.HandleFunc("/{userId}/roles", ra.GetUserRoles).Methods("GET")
}
