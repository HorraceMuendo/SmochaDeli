package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AuthBridge(c *fiber.Ctx) error {

	fmt.Println("in middleware")
	c.Next()
	return c.SendStatus(fiber.StatusOK)
}
