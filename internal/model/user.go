package model

type User struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password,omitempty"`
	CurrentLevel int    `json:"current_level"`
}

type CreateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
