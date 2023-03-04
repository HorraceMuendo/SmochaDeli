package database

import (
	customers "SmochaDeliveryApp/Customers"
	riders "SmochaDeliveryApp/Riders"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	var err error

	DNS := "host=localhost user=postgres password= dbname=smochadeliveryapp port=5432  sslmode=disable"

	DB, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("connection unsuccesful")
	}

	// call the structs
	//fmt.Println("Automigration succesful")
	DB.AutoMigrate(&customers.CustomerDetails{}, &riders.RiderDetails{})
	fmt.Println("Automigration succesful")
}
