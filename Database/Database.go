package database

import (
	customers "SmochaDeliveryApp/Customers"
	riders "SmochaDeliveryApp/Riders"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Conn() *gorm.DB {
	DNS := "host=localhost user=postgres password= dbname=smochadeliveryapp port=5432  sslmode=disabled"

	db, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("connection unsuccesful")
	}
	// call the structs
	err = db.AutoMigrate(&customers.CustomerDetails{}, &riders.RiderDetails{})
	if err != nil {
		log.Fatal("Automigration failed")
	}
	return db
}
