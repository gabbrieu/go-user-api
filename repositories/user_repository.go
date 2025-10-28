package repository

import (
	"context"
	"errors"
	"fmt"
	"user-api/entities"
	"user-api/ports"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{DB: db}
}

func (repository *userRepository) Create(ctx context.Context, dto ports.CreateUserDto) (*entities.User, error) {
	user := &entities.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	if err := repository.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *userRepository) Update(ctx context.Context, id uint, dto ports.UpdateUserDto) (*entities.User, error) {
	updates := map[string]any{}

	if dto.Name != nil {
		updates["name"] = *dto.Name
	}

	tx := repository.WithContext(ctx).Model(&entities.User{}).Where("id = ?", id).Updates(updates).Clauses(clause.Returning{})

	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var updated entities.User
	if err := tx.Scan(&updated).Error; err != nil {
		if err := repository.WithContext(ctx).Where("id = ?", id).First(&updated).Error; err != nil {
			return nil, err
		}
	}

	return &updated, nil
}

func (repository *userRepository) GetAll(ctx context.Context) ([]entities.User, error) {
	var users []entities.User

	if err := repository.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *userRepository) GetOne(ctx context.Context, id uint) (*entities.User, error) {
	var user entities.User

	err := repository.WithContext(ctx).Unscoped().Where(&entities.User{Id: id}).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}

		return nil, err
	}

	return &user, nil
}
