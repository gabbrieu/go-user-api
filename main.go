package main

import (
	"log"
	"user-api/configuration"
	controller "user-api/controllers"
	"user-api/handlers"
	repository "user-api/repositories"
	service "user-api/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := configuration.NewDatabase()
	configuration.Migrate(db)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)

	app := fiber.New(fiber.Config{
		AppName:      "Go User API",
		ErrorHandler: handlers.ErrorHandler,
	})
	userController.Route(app)

	log.Fatal(app.Listen(":3000"))
}
