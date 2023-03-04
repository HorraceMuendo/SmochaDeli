package handlers

import (
	database "SmochaDeliveryApp/Database"
	riders "SmochaDeliveryApp/Riders"

	"github.com/gofiber/fiber/v2"
)

func GetRider(c *fiber.Ctx) error {
	Rdb := database.DB
	var riderDetails []riders.RiderDetails
	Rdb.Find(&riderDetails)
	return c.Status(200).JSON(riderDetails)
}

func GetRiderById(c *fiber.Ctx) error {
	Rdb := database.DB
	id := c.Params("id")
	var riderDetail []riders.RiderDetails
	match := Rdb.Find(&riderDetail, id)

	if match.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&riderDetail)

}
func CreateRider(c *fiber.Ctx) error {
	Rdb := database.DB
	rider := new(riders.RiderDetails)
	if err := c.BodyParser(rider); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	Rdb.Create(rider)
	return c.Status(200).JSON(rider)

}

func UpdateRider(c *fiber.Ctx) error {
	Rdb := database.DB
	rider := new(riders.RiderDetails)
	id := c.Params("id")
	if err := c.BodyParser(rider); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	Rdb.Where("id=?", id).Updates(&rider)
	return c.Status(200).JSON(rider)

}
func DeleteRider(c *fiber.Ctx) error {
	Rdb := database.DB
	var rider riders.RiderDetails
	id := c.Params("id")
	delete := Rdb.Delete(&rider, id)

	if delete.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
