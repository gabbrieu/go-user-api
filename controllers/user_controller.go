package ctl

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

func (ctl UserController) Route(app *fiber.App) {
	app.Get("/users/:id", ctl.GetOne)
	app.Get("/users", ctl.GetAll)
	app.Post("/users", ctl.Create)
	app.Patch("/users/:id", ctl.Update)
}

func (ctl UserController) Create(c *fiber.Ctx) error {
	var dto ports.CreateUserDto

	if err := common.BindAndValidate(c, dto); err != nil {
		return err
	}

	createdUser, err := ctl.UserService.Create(c.Context(), dto)
	if err != nil {
		return common.MapServiceError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

func (ctl UserController) Update(c *fiber.Ctx) error {
	user, err := getOneUser(c, ctl)
	if err != nil {
		return err
	}

	var dto ports.UpdateUserDto

	if err := common.BindAndValidate(c, dto); err != nil {
		return err
	}

	updatedUser, err := ctl.UserService.Update(c.Context(), user.Id, dto)
	if err != nil {
		return common.MapServiceError(err)
	}

	return c.Status(fiber.StatusOK).JSON(updatedUser)
}

func (ctl UserController) GetAll(c *fiber.Ctx) error {
	users, err := ctl.UserService.GetAll(c.Context())
	if err != nil {
		return common.MapServiceError(err)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (ctl UserController) GetOne(c *fiber.Ctx) error {
	user, err := getOneUser(c, ctl)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func getOneUser(ctx *fiber.Ctx, ctl UserController) (*entities.User, error) {
	id := ctx.Params("id")

	idInt, err := strconv.ParseUint(id, 10, 64)
	exception.FatalLogging(err, fmt.Sprintf("Id not converted to unsigned integer: %s", err))

	user, err := ctl.UserService.GetOne(ctx.Context(), uint(idInt))

	if err != nil {
		return nil, common.MapServiceError(err)
	}

	return user, nil
}
