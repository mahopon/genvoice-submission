package service

import (
	"backend/internal/model/user"
	db "backend/internal/repository"
	"backend/internal/util"
	"errors"
	"strings"

	"github.com/google/uuid"
)

type UserService interface {
	LoginUser(user model.LoginUserRequest) (*model.User, error)
	RegisterUser(user model.CreateUserRequest) error
	UpdateUser(id string, user model.UpdateUserRequest) error
	DeleteUser(id string) error
	GetUser(id string) *model.GetUserResponse
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) LoginUser(user model.LoginUserRequest) (*model.User, error) {
	dbUser, err := db.GetUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}
	salt, hash := util.SplitPasswordSalt(dbUser.Password)

	hashInput, err := util.GenerateFromPasswordWithSalt(user.Password, salt)
	if err != nil {
		return nil, err
	}

	if hashInput != hash {
		return nil, errors.New("invalid credentials")
	}

	return dbUser, nil
}

func (s *userService) RegisterUser(user model.CreateUserRequest) error {
	hash, salt, err := util.GenerateFromPassword(user.Password)
	if err != nil {
		return err
	}
	builder := strings.Builder{}
	builder.WriteString(salt)
	builder.WriteString(":")
	builder.WriteString(hash)
	newUser := &model.User{
		Name:     user.Name,
		Username: user.Username,
		Password: builder.String(),
		Role:     "ADMIN",
	}
	err = db.CreateUser(newUser)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) UpdateUser(id string, updatedUser model.UpdateUserRequest) error {
	return nil
}

func (s *userService) DeleteUser(id string) error {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	err = db.DeleteUser(parsedUUID)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetUser(id string) *model.GetUserResponse {
	return &model.GetUserResponse{Name: "AHHH"}
}
