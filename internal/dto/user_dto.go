package dto

import "github.com/oxtx/go-rest-api/internal/model"

type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required,min=2"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func MapUserToResponse(u *model.User) *UserResponse {
	return &UserResponse{ID: u.ID, Email: u.Email, Name: u.Name}
}
