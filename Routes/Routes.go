package routes

import (
	handlers "SmochaDeliveryApp/Handlers"
	middleware "SmochaDeliveryApp/Middleware"
	transactions "SmochaDeliveryApp/Transactions"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Routes() {
	app := fiber.New(fiber.Config{
		//Prefork: true,
		AppName: "SmochaDeliveryApp",
	})

	api := app.Group("/api")

	//customers endpoints
	customer := api.Group("/customers")
	customer.Post("/signup", handlers.SignUpCustomer)
	customer.Post("/login", handlers.LoginCustomer)
	customer.Get("/validate", middleware.CustomerAuthBridge, handlers.ValidateCustomer)

	//riders endpoints
	rider := api.Group("/riders")
	rider.Post("/signup", handlers.SignUpRider)
	rider.Post("/login", handlers.LoginRider)
	rider.Get("/validate", middleware.RiderAuthBridge, handlers.ValidateRider)

	//transactons endpoints
	transaction := api.Group("/transactions")
	transaction.Post("/paybill", transactions.DarajaApi)
	fmt.Println("starting server at port 3000")
	log.Fatal(app.Listen(os.Getenv("PORT")))
}
