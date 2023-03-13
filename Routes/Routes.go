package routes

import (
	handlers "SmochaDeliveryApp/Handlers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Routes() {
	app := fiber.New(fiber.Config{
		Prefork: true,
		AppName: "SmochaDeliveryApp",
	})

	api := app.Group("/api")
	//customers endpoints
	customer := api.Group("/customers")
	customer.Post("/signup", handlers.SignUp)
	customer.Post("/login/:Email", handlers.Login)
	//riders endpoints
	//riders := api.Group("/riders")

	fmt.Println("starting server at port 3000")
	log.Fatal(app.Listen(":3000"))
}

// customer.Get("/getCustomers", handlers.GetCustomer)
// customer.Get("/getCustomerById", handlers.GetCustomerById)
// customer.Post("/createCustomer", handlers.GetCustomer)
// customer.Put("/updateCustomer", handlers.GetCustomer)
// customer.Delete("/deleteCustomer", handlers.GetCustomer)

// riders.Get("/getRiders", handlers.GetRider)
// riders.Get("/getRiderById", handlers.GetRiderById)
// riders.Post("/createRider", handlers.CreateRider)
// riders.Put("/updateRider", handlers.UpdateRider)
// riders.Delete("/deleteRider", handlers.DeleteRider)
