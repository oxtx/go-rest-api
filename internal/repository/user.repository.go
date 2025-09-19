package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/oxtx/go-rest-api/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, email, name string) (*model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, email, name string) (*model.User, error) {
	id := uuid.New().String()
	row := r.db.QueryRowContext(ctx,
		`INSERT INTO users (id,email,name) VALUES ($1,$2,$3) RETURNING created_at`,
		id, email, name)

	var createdAt sql.NullTime
	if err := row.Scan(&createdAt); err != nil {
		return nil, err
	}
	return &model.User{ID: id, Email: email, Name: name, CreatedAt: createdAt.Time}, nil
}

func (r *userRepo) GetByID(ctx context.Context, id string) (*model.User, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT id,email,name,created_at FROM users WHERE id=$1`, id)
	var u model.User
	if err := row.Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}
