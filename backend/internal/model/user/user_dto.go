package model

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID   uint   `son:"id"`
	Name string `json:"name"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
