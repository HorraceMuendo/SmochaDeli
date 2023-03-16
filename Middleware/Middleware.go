package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Middleware(c *fiber.Ctx) error {

	fmt.Println("in middleware")
	return c.JSON("okay")
}
