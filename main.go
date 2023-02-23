package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
		AppName: "SmochaDeliveryApp",
	})
	// connection string
	//fmt.Println("starting server at port 3000")
	log.Fatal(app.Listen(":3000"))

}
