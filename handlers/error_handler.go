package handlers

import "github.com/gofiber/fiber/v2"

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := "Erro interno"

	if fe, ok := err.(*fiber.Error); ok {
		code = fe.Code
		msg = fe.Message
	} else if err != nil {
		msg = err.Error()
	}

	return c.Status(code).JSON(fiber.Map{
		"message": msg,
	})
}
