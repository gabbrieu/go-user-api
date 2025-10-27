package ports

import (
	"context"
	"user-api/entities"
)

type UserService interface {
	Create(ctx context.Context, user CreateUserDto) (*entities.User, error)
	GetOne(ctx context.Context, id uint) (*entities.User, error)
	GetAll(ctx context.Context) ([]entities.User, error)
}

type UserRepository interface {
	Create(ctx context.Context, user CreateUserDto) (*entities.User, error)
	GetOne(ctx context.Context, id uint) (*entities.User, error)
	GetAll(ctx context.Context) ([]entities.User, error)
}

type CreateUserDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
