package models

// User represents a user in the directory system
type User struct {
	ID       string `json:"id"`        		// Unique identifier (e.g., UI000000)
	Email    string `json:"email"`     		// Email address (e.g., user@company.onmicrosoft.com)
	Name     string `json:"name"`      		// Display name (e.g., "John Doe")
	Roles    []Role `json:"roles"`     		// Assigned roles (users can have multiple roles)
	GroupIDs []string `json:"group_ids"` 	// IDs of groups the user belongs to
}

// Group represents a group in the directory system
type Group struct {
	ID          string   `json:"id"`          // Unique group identifier
	Name        string   `json:"name"`        // Group display name
	Description string   `json:"description"` // Group description/purpose
	Members     []string `json:"members"`     // List of member user IDs
}

// Role represents a role with associated permissions
type Role struct {
	ID          string   `json:"id"`          // Unique role identifier
	Name        string   `json:"name"`        // Role display name
	Description string   `json:"description"` // Role description
	Groups      []string `json:"groups"`      // List of group IDs associated with this role
}