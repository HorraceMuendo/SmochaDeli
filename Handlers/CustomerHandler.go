package handlers

import (
	database "SmochaDeliveryApp/Database"
	"SmochaDeliveryApp/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {

	var CustomerSignup model.CustomerDetails

	hash, err := bcrypt.GenerateFromPassword([]byte(CustomerSignup.Password), 10)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)

	}
	CustomerSignup = model.CustomerDetails{
		Firstname: CustomerSignup.Firstname,
		Lastname:  CustomerSignup.Lastname,
		Email:     CustomerSignup.Email,
		Phone:     CustomerSignup.Phone,
		Location:  CustomerSignup.Location,
		Password:  string(hash),
	}
	addUser := database.Db.Create(&CustomerSignup)
	if addUser.Error != nil {
		c.Status(500).JSON("failed to create user")
	}

	return c.SendStatus(fiber.StatusOK)
}
