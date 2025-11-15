package handlers

import (
	"fmt"
	"github.com/lotusatx/lotus-directory-engine-backend/models"
	"gorm.io/gorm"
)

// CreateGroup creates a new group in the database
func CreateGroup(db *gorm.DB, group *models.Group) error {
	result := db.Create(group)
	if result.Error != nil {
		return fmt.Errorf("failed to create group: %w", result.Error)
	}
	return nil
}

// GetGroupByID retrieves a group by its ID
func GetGroupByID(db *gorm.DB, groupID string) (*models.Group, error) {
	var group models.Group
	result := db.Where("id = ?", groupID).First(&group)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("group not found: %s", groupID)
		}
		return nil, fmt.Errorf("failed to get group: %w", result.Error)
	}
	return &group, nil
}

// GetAllGroups retrieves all groups
func GetAllGroups(db *gorm.DB) ([]models.Group, error) {
	var groups []models.Group
	result := db.Order("name").Find(&groups)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to query groups: %w", result.Error)
	}
	return groups, nil
}

// UpdateGroup updates an existing group
func UpdateGroup(db *gorm.DB, group *models.Group) error {
	result := db.Save(group)
	if result.Error != nil {
		return fmt.Errorf("failed to update group: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("group not found: %s", group.ID)
	}
	return nil
}

// DeleteGroup deletes a group by ID
func DeleteGroup(db *gorm.DB, groupID string) error {
	result := db.Delete(&models.Group{}, "id = ?", groupID)
	if result.Error != nil {
		return fmt.Errorf("failed to delete group: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("group not found: %s", groupID)
	}
	return nil
}

// AddUserToGroup adds a single user to a group
func AddUserToGroup(db *gorm.DB, groupID string, userID string) error {
	group, err := GetGroupByID(db, groupID)
	if err != nil {
		return err
	}

	// Check if user already exists in the group
	for _, member := range group.Members {
		if member == userID {
			return fmt.Errorf("user %s is already a member of group %s", userID, groupID)
		}
	}

	// Add user to members list
	group.Members = append(group.Members, userID)
	
	result := db.Save(group)
	if result.Error != nil {
		return fmt.Errorf("failed to add user to group: %w", result.Error)
	}
	return nil
}

// AddUsersToGroup adds multiple users to a group
func AddUsersToGroup(db *gorm.DB, groupID string, userIDs []string) error {
	group, err := GetGroupByID(db, groupID)
	if err != nil {
		return err
	}

	// Create a map of existing members for quick lookup
	existingMembers := make(map[string]bool)
	for _, member := range group.Members {
		existingMembers[member] = true
	}

	// Add only new users
	addedCount := 0
	for _, userID := range userIDs {
		if !existingMembers[userID] {
			group.Members = append(group.Members, userID)
			existingMembers[userID] = true
			addedCount++
		}
	}

	if addedCount == 0 {
		return fmt.Errorf("all specified users are already members of the group")
	}

	result := db.Save(group)
	if result.Error != nil {
		return fmt.Errorf("failed to add users to group: %w", result.Error)
	}
	return nil
}

// RemoveUserFromGroup removes a single user from a group
func RemoveUserFromGroup(db *gorm.DB, groupID string, userID string) error {
	group, err := GetGroupByID(db, groupID)
	if err != nil {
		return err
	}

	// Find and remove the user
	found := false
	newMembers := make([]string, 0, len(group.Members))
	for _, member := range group.Members {
		if member != userID {
			newMembers = append(newMembers, member)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("user %s is not a member of group %s", userID, groupID)
	}

	group.Members = newMembers
	result := db.Save(group)
	if result.Error != nil {
		return fmt.Errorf("failed to remove user from group: %w", result.Error)
	}
	return nil
}

// RemoveUsersFromGroup removes multiple users from a group
func RemoveUsersFromGroup(db *gorm.DB, groupID string, userIDs []string) error {
	group, err := GetGroupByID(db, groupID)
	if err != nil {
		return err
	}

	// Create a map of users to remove
	usersToRemove := make(map[string]bool)
	for _, userID := range userIDs {
		usersToRemove[userID] = true
	}

	// Filter out the users to remove
	newMembers := make([]string, 0, len(group.Members))
	removedCount := 0
	for _, member := range group.Members {
		if !usersToRemove[member] {
			newMembers = append(newMembers, member)
		} else {
			removedCount++
		}
	}

	if removedCount == 0 {
		return fmt.Errorf("none of the specified users are members of the group")
	}

	group.Members = newMembers
	result := db.Save(group)
	if result.Error != nil {
		return fmt.Errorf("failed to remove users from group: %w", result.Error)
	}
	return nil
}

// GetGroupMembers retrieves all user IDs that are members of a group
func GetGroupMembers(db *gorm.DB, groupID string) ([]string, error) {
	group, err := GetGroupByID(db, groupID)
	if err != nil {
		return nil, err
	}
	return group.Members, nil
}

// GetUserGroups retrieves all groups that a user is a member of
func GetUserGroups(db *gorm.DB, userID string) ([]models.Group, error) {
	var groups []models.Group
	// This assumes Members is stored as a JSON array or similar
	// The exact query depends on your database schema
	result := db.Where("? = ANY(members)", userID).Find(&groups)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get user groups: %w", result.Error)
	}
	return groups, nil
}
