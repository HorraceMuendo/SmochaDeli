package handlers

import (
	database "SmochaDeliveryApp/Database"
	"SmochaDeliveryApp/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) {

	var CustomerSignup model.CustomerDetails

	hash, err := bcrypt.GenerateFromPassword([]byte(Customer.Password), 10)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return
	}
	CustomerSignup = model.CustomerDetails{
		Firstname: CustomerSignup.Firstname,
		Lastname:  CustomerSignup.Lastname,
		Email:     CustomerSignup.Email,
		Phone:     CustomerSignup.Phone,
		Location:  CustomerSignup.Location,
		Password:  string(hash),
	}
	database.Db.Create(&CustomerSignup)

	return
}
