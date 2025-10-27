package main

import (
	"log"
	"user-api/configuration"
	controller "user-api/controllers"
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

	app := fiber.New()
	userController.Route(app)

	log.Fatal(app.Listen(":3000"))
}
