package database

import (
	customers "SmochaDeliveryApp/Customers"
	riders "SmochaDeliveryApp/Riders"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Conn() {
	var err error

	DNS := "host=localhost user=postgres password=muendo dbname=smochadeliveryapp port=5432  sslmode=disable"

	Db, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("connection unsuccesful")
	}

	Db.AutoMigrate(&customers.CustomerDetails{}, &riders.RiderDetails{})
	fmt.Println("Automigration succesful")
}
