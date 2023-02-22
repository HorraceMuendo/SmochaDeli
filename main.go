package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/prefork"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New(fiber.Config{
		prefork : true,
		appName:"SmochaDelivery"
	})

	db,err:= gorm.Open(postgres.New(postgres.Config{
		
	}))

	app.Listen(":3000")

}
