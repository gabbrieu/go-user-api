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

func (service *userService) Create(ctx context.Context, createUserDTO ports.CreateUserDto) (*entities.User, error) {
	return service.UserRepository.Create(ctx, createUserDTO)
}

func (service *userService) GetOne(ctx context.Context, id uint) (*entities.User, error) {
	return service.UserRepository.GetOne(ctx, id)
}

func (service *userService) GetAll(ctx context.Context) ([]entities.User, error) {
	return service.UserRepository.GetAll(ctx)
}
