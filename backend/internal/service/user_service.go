package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/util"
	"errors"

	"github.com/google/uuid"
)

type UserService interface {
	LoginUser(user model.LoginUserRequest) (*model.User, error)
	RegisterUser(user model.CreateUserRequest) error
	UpdateUser(id string, user model.UpdateUserPasswordRequest) error
	UpdateWholeUser(userId string, request model.UpdateUserRequest) error
	DeleteUser(id string) error
	GetUser(id string) (*model.GetUserResponse, error)
	GetAllUser() (*[]model.User, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) LoginUser(user model.LoginUserRequest) (*model.User, error) {
	dbUser, err := repository.GetUserByUsername(user.Username)
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

	var role string
	if user.Role == "" {
		role = "USER"
	} else {
		role = user.Role
	}

	password := salt + ":" + hash
	newUser := &model.User{
		Name:     user.Name,
		Username: user.Username,
		Password: password,
		Role:     role,
	}

	return repository.CreateUser(newUser)
}

func (s *userService) UpdateUser(id string, update model.UpdateUserPasswordRequest) error {
	userID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	existingUser, err := repository.GetUserByID(userID)
	if err != nil {
		return err
	}

	salt, hash := util.SplitPasswordSalt(existingUser.Password)
	inputPassword, _ := util.GenerateFromPasswordWithSalt(update.CurrentPassword, salt)

	if hash != inputPassword {
		return errors.New("current password wrong")
	}

	newPassword, salt, _ := util.GenerateFromPassword(update.NewPassword)
	stitchedNewPassword := salt + ":" + newPassword

	err = repository.UpdateUser(existingUser.ID, stitchedNewPassword)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(id string) error {
	userID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return repository.DeleteUser(userID)
}

func (s *userService) GetUser(id string) (*model.GetUserResponse, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	user, err := repository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return &model.GetUserResponse{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (s *userService) GetAllUser() (*[]model.User, error) {
	var users *[]model.User
	users, err := repository.GetAllUser()
	if err != nil {
		return nil, errors.New("internal error")
	}
	return users, nil
}

func (s *userService) UpdateWholeUser(userId string, request model.UpdateUserRequest) error {
	parsedID, _ := uuid.Parse(userId)

	if request.Password != "" {
		hash, salt, _ := util.GenerateFromPassword(request.Password)
		request.Password = salt + ":" + hash
	}

	err := repository.UpdateWholeUser(parsedID, &request)
	if err != nil {
		return err
	}
	return nil
}
