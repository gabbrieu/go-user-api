package exception

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func FatalLogging(err error, msg string) {
	if err != nil {
		log.Fatalf("%s", msg)
	}
}

func ParseError(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{
		"message": fmt.Sprintf("Parse error - %s", msg),
	})
}
