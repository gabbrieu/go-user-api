package controller

import (
	"fmt"
	"strconv"
	"user-api/common"
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
	app.Get("/users", controller.GetAll)
	app.Post("/users", controller.Create)
}

func (controller UserController) Create(c *fiber.Ctx) error {
	var dto ports.CreateUserDto

	if err := c.BodyParser(&dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}

	if err := common.Validate(dto); err != nil {
		if vErr, ok := err.((*common.ValidationError)); ok {
			return c.Status(fiber.StatusBadRequest).JSON(vErr)
		}
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	createdUser, err := controller.UserService.Create(c.Context(), dto)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

func (controller UserController) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.ParseUint(id, 10, 64)
	exception.FatalLogging(err, fmt.Sprintf("id not converted to unsigned integer: %s", err))

	user, err := controller.UserService.GetOne(c.Context(), uint(idInt))

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (controller UserController) GetAll(c *fiber.Ctx) error {
	users, err := controller.UserService.GetAll(c.Context())

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
