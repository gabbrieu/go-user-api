package common

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func MapServiceError(err error) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return fiber.NewError(fiber.StatusNotFound, "Resource not found")
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return fiber.NewError(fiber.StatusConflict, "Duplicated resource")
	default:
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
}
