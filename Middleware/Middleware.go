package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthBridge(c *fiber.Ctx) error {
	//get the cookie off the request body
	//decode and validate
	//check the expiration
	//find user with subject
	//attach to the req body
	//continue
	tokenstr := c.Cookies("Authorization")

	c.Next()
	return c.SendStatus(fiber.StatusOK)
}
