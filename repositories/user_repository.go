package repository

import (
	"context"
	"fmt"
	"user-api/entities"
	"user-api/exception"
	"user-api/ports"

	"gorm.io/gorm"
)

type userRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{DB: db}
}

func (repository *userRepository) Create(ctx context.Context, user ports.CreateUserDto) entities.User {
	var createdUser entities.User

	err := repository.DB.WithContext(ctx).Create(&createdUser).Error

	exception.FatalLogging(err, fmt.Sprintf("error while creating the user: %s", err))

	return createdUser
}

func (repository *userRepository) GetAll(ctx context.Context) []entities.User {
	var users []entities.User

	repository.DB.WithContext(ctx).Find(&users)
	return users
}

func (repository *userRepository) GetOne(ctx context.Context, id uint) (entities.User, error) {
	var user entities.User

	result := repository.DB.WithContext(ctx).Unscoped().Where(&entities.User{Id: id}).First(&user)

	if result.RowsAffected == 0 {
		return entities.User{}, fmt.Errorf("user with id %d not found", id)
	}

	return user, nil
}
