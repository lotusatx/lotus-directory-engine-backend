package handlers

import (
	"fmt"
	"github.com/lotusatx/lotus-directory-engine-backend/models"
	"gorm.io/gorm"
)

// CreateRole creates a new role in the database
func CreateRole(db *gorm.DB, role *models.Role) error {
	result := db.Create(role)
	if result.Error != nil {
		return fmt.Errorf("failed to create role: %w", result.Error)
	}
	return nil
}

// GetRoleByID retrieves a role by its ID
func GetRoleByID(db *gorm.DB, roleID string) (*models.Role, error) {
	var role models.Role
	result := db.Where("id = ?", roleID).First(&role)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("role not found: %s", roleID)
		}
		return nil, fmt.Errorf("failed to get role: %w", result.Error)
	}
	return &role, nil
}

// GetAllRoles retrieves all roles
func GetAllRoles(db *gorm.DB) ([]models.Role, error) {
	var roles []models.Role
	result := db.Order("name").Find(&roles)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to query roles: %w", result.Error)
	}
	return roles, nil
}

// UpdateRole updates an existing role
func UpdateRole(db *gorm.DB, role *models.Role) error {
	result := db.Save(role)
	if result.Error != nil {
		return fmt.Errorf("failed to update role: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("role not found: %s", role.ID)
	}
	return nil
}

// DeleteRole deletes a role by ID
func DeleteRole(db *gorm.DB, roleID string) error {
	// Note: Consider checking if role is assigned to users before deletion
	result := db.Delete(&models.Role{}, "id = ?", roleID)
	if result.Error != nil {
		return fmt.Errorf("failed to delete role: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("role not found: %s", roleID)
	}
	return nil
}

// AddGroupToRole adds a single group to a role
func AddGroupToRole(db *gorm.DB, roleID string, groupID string) error {
	role, err := GetRoleByID(db, roleID)
	if err != nil {
		return err
	}

	// Check if group already exists in the role
	for _, group := range role.Groups {
		if group == groupID {
			return fmt.Errorf("group %s is already associated with role %s", groupID, roleID)
		}
	}

	// Add group to role's groups list
	role.Groups = append(role.Groups, groupID)
	
	result := db.Save(role)
	if result.Error != nil {
		return fmt.Errorf("failed to add group to role: %w", result.Error)
	}
	return nil
}

// AddGroupsToRole adds multiple groups to a role
func AddGroupsToRole(db *gorm.DB, roleID string, groupIDs []string) error {
	role, err := GetRoleByID(db, roleID)
	if err != nil {
		return err
	}

	// Create a map of existing groups for quick lookup
	existingGroups := make(map[string]bool)
	for _, group := range role.Groups {
		existingGroups[group] = true
	}

	// Add only new groups
	addedCount := 0
	for _, groupID := range groupIDs {
		if !existingGroups[groupID] {
			role.Groups = append(role.Groups, groupID)
			existingGroups[groupID] = true
			addedCount++
		}
	}

	if addedCount == 0 {
		return fmt.Errorf("all specified groups are already associated with the role")
	}

	result := db.Save(role)
	if result.Error != nil {
		return fmt.Errorf("failed to add groups to role: %w", result.Error)
	}
	return nil
}

// RemoveGroupFromRole removes a single group from a role
func RemoveGroupFromRole(db *gorm.DB, roleID string, groupID string) error {
	role, err := GetRoleByID(db, roleID)
	if err != nil {
		return err
	}

	// Find and remove the group
	found := false
	newGroups := make([]string, 0, len(role.Groups))
	for _, group := range role.Groups {
		if group != groupID {
			newGroups = append(newGroups, group)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("group %s is not associated with role %s", groupID, roleID)
	}

	role.Groups = newGroups
	result := db.Save(role)
	if result.Error != nil {
		return fmt.Errorf("failed to remove group from role: %w", result.Error)
	}
	return nil
}

// RemoveGroupsFromRole removes multiple groups from a role
func RemoveGroupsFromRole(db *gorm.DB, roleID string, groupIDs []string) error {
	role, err := GetRoleByID(db, roleID)
	if err != nil {
		return err
	}

	// Create a map of groups to remove
	groupsToRemove := make(map[string]bool)
	for _, groupID := range groupIDs {
		groupsToRemove[groupID] = true
	}

	// Filter out the groups to remove
	newGroups := make([]string, 0, len(role.Groups))
	removedCount := 0
	for _, group := range role.Groups {
		if !groupsToRemove[group] {
			newGroups = append(newGroups, group)
		} else {
			removedCount++
		}
	}

	if removedCount == 0 {
		return fmt.Errorf("none of the specified groups are associated with the role")
	}

	role.Groups = newGroups
	result := db.Save(role)
	if result.Error != nil {
		return fmt.Errorf("failed to remove groups from role: %w", result.Error)
	}
	return nil
}

// AssignRoleToUser assigns a single role to a user
func AssignRoleToUser(db *gorm.DB, userID string, roleID string) error {
	user, err := GetUserByID(db, userID)
	if err != nil {
		return err
	}

	// Verify the role exists
	_, err = GetRoleByID(db, roleID)
	if err != nil {
		return err
	}

	// Check if user already has this role
	for _, role := range user.Roles {
		if role.ID == roleID {
			return fmt.Errorf("user %s already has role %s", userID, roleID)
		}
	}

	// Fetch the complete role and add it to user
	roleToAdd, err := GetRoleByID(db, roleID)
	if err != nil {
		return err
	}

	user.Roles = append(user.Roles, *roleToAdd)
	
	result := db.Save(user)
	if result.Error != nil {
		return fmt.Errorf("failed to assign role to user: %w", result.Error)
	}
	return nil
}

// AssignRolesToUser assigns multiple roles to a user
func AssignRolesToUser(db *gorm.DB, userID string, roleIDs []string) error {
	user, err := GetUserByID(db, userID)
	if err != nil {
		return err
	}

	// Create a map of existing roles for quick lookup
	existingRoles := make(map[string]bool)
	for _, role := range user.Roles {
		existingRoles[role.ID] = true
	}

	// Fetch and add only new roles
	addedCount := 0
	for _, roleID := range roleIDs {
		if !existingRoles[roleID] {
			role, err := GetRoleByID(db, roleID)
			if err != nil {
				return fmt.Errorf("failed to get role %s: %w", roleID, err)
			}
			user.Roles = append(user.Roles, *role)
			existingRoles[roleID] = true
			addedCount++
		}
	}

	if addedCount == 0 {
		return fmt.Errorf("user already has all specified roles")
	}

	result := db.Save(user)
	if result.Error != nil {
		return fmt.Errorf("failed to assign roles to user: %w", result.Error)
	}
	return nil
}

// RemoveRoleFromUser removes a single role from a user
func RemoveRoleFromUser(db *gorm.DB, userID string, roleID string) error {
	user, err := GetUserByID(db, userID)
	if err != nil {
		return err
	}

	// Find and remove the role
	found := false
	newRoles := make([]models.Role, 0, len(user.Roles))
	for _, role := range user.Roles {
		if role.ID != roleID {
			newRoles = append(newRoles, role)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("user %s does not have role %s", userID, roleID)
	}

	user.Roles = newRoles
	result := db.Save(user)
	if result.Error != nil {
		return fmt.Errorf("failed to remove role from user: %w", result.Error)
	}
	return nil
}

// RemoveRolesFromUser removes multiple roles from a user
func RemoveRolesFromUser(db *gorm.DB, userID string, roleIDs []string) error {
	user, err := GetUserByID(db, userID)
	if err != nil {
		return err
	}

	// Create a map of roles to remove
	rolesToRemove := make(map[string]bool)
	for _, roleID := range roleIDs {
		rolesToRemove[roleID] = true
	}

	// Filter out the roles to remove
	newRoles := make([]models.Role, 0, len(user.Roles))
	removedCount := 0
	for _, role := range user.Roles {
		if !rolesToRemove[role.ID] {
			newRoles = append(newRoles, role)
		} else {
			removedCount++
		}
	}

	if removedCount == 0 {
		return fmt.Errorf("user does not have any of the specified roles")
	}

	user.Roles = newRoles
	result := db.Save(user)
	if result.Error != nil {
		return fmt.Errorf("failed to remove roles from user: %w", result.Error)
	}
	return nil
}

// BulkAssignRoleToUsers assigns a role to multiple users
func BulkAssignRoleToUsers(db *gorm.DB, roleID string, userIDs []string) error {
	// Verify the role exists
	role, err := GetRoleByID(db, roleID)
	if err != nil {
		return err
	}

	successCount := 0
	var errors []string

	for _, userID := range userIDs {
		err := AssignRoleToUser(db, userID, roleID)
		if err != nil {
			errors = append(errors, fmt.Sprintf("user %s: %v", userID, err))
		} else {
			successCount++
		}
	}

	if successCount == 0 {
		return fmt.Errorf("failed to assign role %s to any users: %v", role.Name, errors)
	}

	if len(errors) > 0 {
		return fmt.Errorf("assigned role to %d users, but encountered errors: %v", successCount, errors)
	}

	return nil
}

// BulkRemoveRoleFromUsers removes a role from multiple users
func BulkRemoveRoleFromUsers(db *gorm.DB, roleID string, userIDs []string) error {
	// Verify the role exists
	role, err := GetRoleByID(db, roleID)
	if err != nil {
		return err
	}

	successCount := 0
	var errors []string

	for _, userID := range userIDs {
		err := RemoveRoleFromUser(db, userID, roleID)
		if err != nil {
			errors = append(errors, fmt.Sprintf("user %s: %v", userID, err))
		} else {
			successCount++
		}
	}

	if successCount == 0 {
		return fmt.Errorf("failed to remove role %s from any users: %v", role.Name, errors)
	}

	if len(errors) > 0 {
		return fmt.Errorf("removed role from %d users, but encountered errors: %v", successCount, errors)
	}

	return nil
}

// GetUserRoles retrieves all roles assigned to a user
func GetUserRoles(db *gorm.DB, userID string) ([]models.Role, error) {
	user, err := GetUserByID(db, userID)
	if err != nil {
		return nil, err
	}
	return user.Roles, nil
}

// GetRoleGroups retrieves all groups associated with a role
func GetRoleGroups(db *gorm.DB, roleID string) ([]string, error) {
	role, err := GetRoleByID(db, roleID)
	if err != nil {
		return nil, err
	}
	return role.Groups, nil
}
