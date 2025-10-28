package controller

import (
	"fmt"
	"strconv"
	"user-api/common"
	"user-api/entities"
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
	app.Patch("/users/:id", controller.Update)
}

func (controller UserController) Create(c *fiber.Ctx) error {
	var dto ports.CreateUserDto

	if err := c.BodyParser(&dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := common.Validate(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	createdUser, err := controller.UserService.Create(c.Context(), dto)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

func (controller UserController) Update(c *fiber.Ctx) error {
	user, err := getOneUser(c, controller)

	if err != nil {
		return err
	}

	var dto ports.UpdateUserDto

	if err := c.BodyParser(&dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := common.Validate(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	updatedUser, err := controller.UserService.Update(c.Context(), user.Id, dto)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(updatedUser)
}

func (controller UserController) GetAll(c *fiber.Ctx) error {
	users, err := controller.UserService.GetAll(c.Context())

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (controller UserController) GetOne(c *fiber.Ctx) error {
	user, err := getOneUser(c, controller)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func getOneUser(ctx *fiber.Ctx, controller UserController) (*entities.User, error) {
	id := ctx.Params("id")

	idInt, err := strconv.ParseUint(id, 10, 64)
	exception.FatalLogging(err, fmt.Sprintf("Id not converted to unsigned integer: %s", err))

	user, err := controller.UserService.GetOne(ctx.Context(), uint(idInt))

	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return user, nil
}
