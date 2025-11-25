package ports

import (
	"context"
	"user-api/entities"
)

type UserService interface {
	Create(ctx context.Context, user CreateUserDto) (*entities.User, error)
	GetOne(ctx context.Context, id uint) (*entities.User, error)
	GetAll(ctx context.Context) ([]entities.User, error)
	Update(ctx context.Context, id uint, updateUserDTO UpdateUserDto) (*entities.User, error)
	Delete(ctx context.Context, id uint) error
}

type UserRepository interface {
	Create(ctx context.Context, user CreateUserDto) (*entities.User, error)
	GetOne(ctx context.Context, id uint) (*entities.User, error)
	GetAll(ctx context.Context) ([]entities.User, error)
	Update(ctx context.Context, id uint, updateUserDTO UpdateUserDto) (*entities.User, error)
	Delete(ctx context.Context, id uint) error
}

type CreateUserDto struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type UpdateUserDto struct {
	Name *string `json:"name" validate:"omitempty,min=2"`
}
