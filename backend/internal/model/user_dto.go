package model

import "github.com/google/uuid"

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role,omitempty"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserResponse struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"role"`
}

type UpdateUserPasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type UserResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type GetUserResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
