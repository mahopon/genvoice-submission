package service

import (
	"backend/internal/model/user"
)

type UserService interface {
	LoginUser(user model.LoginUserRequest) (*model.UserResponse, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

// Temp implementation
func (s *userService) LoginUser(user model.LoginUserRequest) (*model.UserResponse, error) {
	return &model.UserResponse{ID: 123, Name: "AHHH"}, nil
}
