package service

import (
	"context"

	"github.com/oxtx/go-rest-api/internal/model"
	"github.com/oxtx/go-rest-api/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, email, name string) (*model.User, error)
	GetUser(ctx context.Context, id string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(ctx context.Context, email, name string) (*model.User, error) {
	return s.repo.Create(ctx, email, name)
}

func (s *userService) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}
