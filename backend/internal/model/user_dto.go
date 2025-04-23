package model

import "github.com/google/uuid"

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type GetUserResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
