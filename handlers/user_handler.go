package handlers

import (
	"fmt"
	"github.com/lotusatx/lotus-directory-engine-backend/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	result := db.Create(user)
	if result.Error != nil {
		return fmt.Errorf("failed to create user: %w", result.Error)
	}
	return nil
}

func GetUserByID(db *gorm.DB, userID string) (*models.User, error) {
	var user models.User
	result := db.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found: %s", userID)
		}
		return nil, fmt.Errorf("failed to get user: %w", result.Error)
	}
	return &user, nil
}

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Order("name").Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to query users: %w", result.Error)
	}
	return users, nil
}

func UpdateUser(db *gorm.DB, user *models.User) error {
	result := db.Save(user)
	if result.Error != nil {
		return fmt.Errorf("failed to update user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found: %s", user.ID)
	}
	return nil
}

func DeleteUser(db *gorm.DB, userID string) error {
	result := db.Delete(&models.User{}, "id = ?", userID)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found: %s", userID)
	}
	return nil
}