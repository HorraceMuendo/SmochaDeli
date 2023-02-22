package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	// connection string
	DNS := "host= user= password= dbname= port=  sslmode=disabled"
	// connecting to the database
	db, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})

	app.Listen(":3000")
	fmt.Println("starting server at port 3000")

}
