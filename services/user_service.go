package service

import (
	"context"
	"user-api/entities"
	"user-api/ports"
)

type userService struct {
	ports.UserRepository
}

func NewUserService(userRepository *ports.UserRepository) ports.UserService {
	return &userService{UserRepository: *userRepository}
}

func (service *userService) Create(ctx context.Context, createUserDTO ports.CreateUserDto) entities.User {

}

func (service *userService) GetOne(ctx context.Context, id uint) entities.User {}

func (service *userService) GetAll(ctx context.Context) []entities.User {}
