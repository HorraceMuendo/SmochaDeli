package handlers

import (
	database "SmochaDeliveryApp/Database"
	"SmochaDeliveryApp/model"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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
			//change the message
			"message": "bad request",
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

	var CustomerLogin model.CustomerDetails
	database.Db.First(&CustomerLogin.Email, "email = ?", body.Email)
	if CustomerLogin.ID == 0 {
		c.Status(400).JSON(fiber.Map{
			"success ?": false,
			"message":   "Invalid Email or Password",
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(CustomerLogin.Password), []byte(body.Password))
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"success ?": false,
			"message":   "password does not match",
		})
	}
	//generating token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"subject": CustomerLogin.ID,
		"expire":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	//signing and encoding
	tokenstr, err := token.SignedString([]byte(os.Getenv("KEY")))
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"success ?": false,
			"message":   "token creaton failure",
		})
	}
	c.Status(200).JSON(tokenstr)
	return c.Status(200).JSON("login succesful...")
}
