package repository

import (
	model "backend/internal/model/user"
	"fmt"

	"github.com/google/uuid"
)

// CreateUser inserts a new user into the database
func CreateUser(user *model.User) error {
	if err := db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// GetUserByUsername fetches a user by their username
func GetUserById(Id int) (*model.User, error) {
	var user model.User
	if err := db.First(&user, "ID = ?", Id).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &user, nil
}

// GetUserByUsername fetches a user by their username
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := db.First(&user, "username = ?", username).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &user, nil
}

// UpdateUser updates a user's details
func UpdateUser(user *model.User) error {
	if err := db.Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// DeleteUser deletes a user by ID
func DeleteUser(id uuid.UUID) error {
	if err := db.Delete(&model.User{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
