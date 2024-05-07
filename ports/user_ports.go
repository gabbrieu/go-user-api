package ports

import (
	"context"
	"user-api/entities"
)

type UserService interface {
	Create(ctx context.Context, user CreateUserDto) entities.User
	GetOne(ctx context.Context, id uint) entities.User
	GetAll(ctx context.Context) []entities.User
}

type UserRepository interface {
	Create(ctx context.Context, user CreateUserDto) entities.User
	GetOne(ctx context.Context, id uint) (entities.User, error)
	GetAll(ctx context.Context) []entities.User
}

type CreateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
