package repository

import (
	"backend/internal/model"
	"fmt"

	"github.com/google/uuid"
)

func CreateUser(user *model.User) error {
	if err := db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func GetUserByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &user, nil
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := db.First(&user, "username = ?", username).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &user, nil
}

func UpdateUser(user *model.User) error {
	if err := db.Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func DeleteUser(id uuid.UUID) error {
	if err := db.Delete(&model.User{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
