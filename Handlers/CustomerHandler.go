package handlers

import (
	database "SmochaDeliveryApp/Database"
	"SmochaDeliveryApp/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {

	var body struct {
		Firstname string
		Lastname  string
		Email     string
		Phone     uint
		Location  string
		Password  string
	}
	err := c.BodyParser(&body)
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"success ?": false,
			"message":   "bad request",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)

	}
	CustomerSignup := model.CustomerDetails{
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Email:     body.Email,
		Phone:     body.Phone,
		Location:  body.Location,
		Password:  string(hash),
	}
	addUser := database.Db.Create(&CustomerSignup)
	if addUser.Error != nil {
		c.Status(500).JSON("failed to create user")
	}

	return c.SendStatus(fiber.StatusOK)
}

func Login(c *fiber.Ctx) error {
	//Email := c.Params("Email")
	var CustomerLogin model.CustomerDetails

	err := c.BodyParser(&CustomerLogin)
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"success ?": false,
			"message":   "failed to bind to struct",
		})
	}

	database.Db.First(&CustomerLogin.Email, "email = ?")
	if CustomerLogin.ID == 0 {
		c.Status(400).JSON(fiber.Map{
			"success ?": false,
			"message":   "Invalid Email or Password",
		})
	}

	return
}
