package common

import "github.com/gofiber/fiber/v2"

func BindAndValidate(c *fiber.Ctx, dto any) error {
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	if err := Validate(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return nil
}
