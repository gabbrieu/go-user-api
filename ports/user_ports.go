package ports

import (
	"context"
	"user-api/entities"
)

type UserService interface {
	Create(ctx context.Context, user entities.User) entities.User
	GetOne(ctx context.Context, id string) entities.User
	GetAll(ctx context.Context) []entities.User
}

type UserRepository interface {
	Create(ctx context.Context, user entities.User) entities.User
	GetOne(ctx context.Context, id string) entities.User
	GetAll(ctx context.Context) []entities.User
}
