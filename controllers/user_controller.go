package controllers

import (
	"user-api/ports"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	ports.UserService
}

func NewUserController(userService *ports.UserService) *UserController {
	return &UserController{UserService: *userService}
}

func (controller UserController) Route(app *fiber.App) {
	app.Get("/users/:id", controller.GetOne)
}

func (controller UserController) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")

	user := controller.UserService.GetOne(c.Context(), id)

	return c.Status(fiber.StatusOK).JSON(user)
}
