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
	app.Delete("/users/:id", ctl.DeleteUser)
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

func getOneUser(c *fiber.Ctx, ctl UserController) (*entities.User, error) {
	id := c.Params("id")

	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, exception.ParseError(c, fmt.Sprintf("Id not converted to unsigned integer: %s", err.Error()))
	}

	user, err := ctl.UserService.GetOne(c.Context(), uint(idInt))

	if err != nil {
		return nil, common.MapServiceError(err)
	}

	return user, nil
}

func (ctl UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return exception.ParseError(c, fmt.Sprintf("Id not converted to unsigned integer: %s", err.Error()))
	}

	if err := ctl.UserService.Delete(c.Context(), uint(idInt)); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
