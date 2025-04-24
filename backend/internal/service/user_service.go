package service

import "backend/internal/model"

type UserService interface {
	LoginUser(user model.LoginUserRequest) (*model.User, error)
	RegisterUser(user model.CreateUserRequest) error
	UpdateUser(id string, user model.UpdateUserPasswordRequest) error
	UpdateWholeUser(userId string, request model.UpdateUserRequest) error
	DeleteUser(id string) error
	GetUser(id string) (*model.GetUserResponse, error)
	GetAllUser() (*[]model.User, error)
}
