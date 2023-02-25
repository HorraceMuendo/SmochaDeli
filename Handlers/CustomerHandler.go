package handlers

import (
	customers "SmochaDeliveryApp/Customers"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

var db *gorm.DB

func Get(c *fiber.Ctx) error {
	var customerDetails []customers.CustomerDetails
	db.Find(&customerDetails)
	return c.Status(200).JSON(customerDetails)
}

func GetId(c *fiber.Ctx) error {
	id := c.Params("id")
	var customerDetail []customers.CustomerDetails
	match := db.Find(&customerDetail, id)
	if match.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.status(200).JSON(&customerDetail)

}
func Create(c *fiber.Ctx) error {
	customer := new(customers.CustomerDetails)
	if err := c.BodyParser(customer); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(customer)
	return c.SendStatus(200).JSON(customer)
}

func Update(c *fiber.Ctx) error {

}
func Delete(c *fiber.Ctx) error {

}
