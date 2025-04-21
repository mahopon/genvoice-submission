package service

import (
	"backend/internal/model/user"
)

type UserService interface {
	LoginUser(user model.LoginUserRequest) *model.UserResponse
	RegisterUser(user model.CreateUserRequest) *model.UserResponse
	UpdateUser(id string, user model.UpdateUserRequest) *model.UserResponse
	DeleteUser(id string) *model.UserResponse
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

// Temp implementation
func (s *userService) LoginUser(user model.LoginUserRequest) *model.UserResponse {
	return &model.UserResponse{Message: "Success"}
}

func (s *userService) RegisterUser(user model.CreateUserRequest) *model.UserResponse {
	// hash, salt, _ := util.GenerateFromPassword(user.Password) // Working implementation
	return &model.UserResponse{Message: "Success"}
}

func (s *userService) UpdateUser(id string, updatedUser model.UpdateUserRequest) *model.UserResponse {
	return &model.UserResponse{Message: "Success"}
}

func (s *userService) DeleteUser(id string) *model.UserResponse {
	// Convert to int and write delete logic
	return &model.UserResponse{Message: "Success"}
}
