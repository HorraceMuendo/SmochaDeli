package handlers

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

var db *gorm.DB

func Get(c *fiber.Ctx) error {
	var customerDetails []CustomerDetails
	db.Find(customerDetails)
	return c.Status(200).JSON(&customerDetails)
}

func GetId() {

}
func Create() {

}

func Update() {

}
func Delete() {

}
