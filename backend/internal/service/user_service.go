package service

import (
	"backend/internal/model"

	"github.com/google/uuid"
)

type UserService interface {
	LoginUser(user *model.LoginUserRequest) (*model.User, error)
	RegisterUser(user *model.CreateUserRequest) error
	UpdateUser(id uuid.UUID, user *model.UpdateUserPasswordRequest) error
	UpdateWholeUser(userId string, request *model.UpdateUserRequest) error
	DeleteUser(id string) error
	GetUser(id uuid.UUID) (*model.GetUserResponse, error)
	GetAllUser() (*[]model.User, error)
}
