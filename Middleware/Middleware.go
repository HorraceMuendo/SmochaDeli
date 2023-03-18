package middleware

import (
	database "SmochaDeliveryApp/Database"
	"SmochaDeliveryApp/model"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthBridge(c *fiber.Ctx) error {
	//get the cookie off the request body
	//decode and validate
	//check the expiration
	//find user with subject
	//attach to the req body
	//continue
	tokenString := c.Cookies("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check expiration
		// if time.Now().Unix() > claims["expire"]{
		// 	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		// 		"message": "unauthorized",
		// 	})
		// }

		// get customer
		var customer model.CustomerDetails
		database.Db.First(&customer, claims["subject"])
		if customer.ID == 0 {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "failed to get user",
			})
		}
		//attach to the request body

		c.Set("customer", customer)

		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	c.Next()
	return c.SendStatus(fiber.StatusOK)
}
