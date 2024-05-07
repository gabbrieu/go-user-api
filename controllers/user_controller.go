package controllers

import (
	"fmt"
	"strconv"
	"user-api/exception"
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

	idInt, err := strconv.ParseUint(id, 10, 64)
	exception.FatalLogging(err, fmt.Sprintf("id not converted to unsigned integer: %s", err))

	user := controller.UserService.GetOne(c.Context(), uint(idInt))

	return c.Status(fiber.StatusOK).JSON(user)
}
