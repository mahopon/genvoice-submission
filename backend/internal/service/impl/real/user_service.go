package real

import (
	"backend/internal/model"
	repo "backend/internal/repository"
	"backend/internal/service"
	"backend/internal/util"
	"errors"

	"github.com/google/uuid"
)

type userService struct{}

func NewUserService() service.UserService {
	return &userService{}
}

func (s *userService) LoginUser(user *model.LoginUserRequest) (*model.User, error) {
	dbUser, err := repo.GetUserByUsername(user.Username)
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

func (s *userService) RegisterUser(user *model.CreateUserRequest) error {
	hash, salt, _ := util.GenerateFromPassword(user.Password)

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

	err := repo.CreateUser(newUser)
	if err != nil {
		return errors.New("user exists")
	}

	return nil
}

func (s *userService) UpdateUser(id uuid.UUID, update *model.UpdateUserPasswordRequest) error {
	existingUser, err := repo.GetUserByID(id)
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

	err = repo.UpdateUser(existingUser.ID, stitchedNewPassword)
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
	return repo.DeleteUser(userID)
}

func (s *userService) GetUser(id uuid.UUID) (*model.GetUserResponse, error) {

	user, err := repo.GetUserByID(id)
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
	users, err := repo.GetAllUser()
	if err != nil {
		return nil, errors.New("internal error")
	}
	return users, nil
}

func (s *userService) UpdateWholeUser(userId string, request *model.UpdateUserRequest) error {
	parsedID, _ := uuid.Parse(userId)

	if request.Password != "" {
		hash, salt, _ := util.GenerateFromPassword(request.Password)
		request.Password = salt + ":" + hash
	}

	err := repo.UpdateWholeUser(parsedID, request)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) CheckAdminRole(userId uuid.UUID) error {
	role, err := repo.GetRole(userId)
	if err != nil {
		return errors.New("invalid user")
	}
	if role != "ADMIN" {
		return errors.New("invalid request")
	}
	return nil
}
