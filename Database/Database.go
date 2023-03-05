package database

import (
	customers "SmochaDeliveryApp/Customers"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Conn() {
	var err error

	DNS := "host=localhost user=postgres password= dbname=smochadeliveryapp port=5432  sslmode=disable"

	Db, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("connection unsuccesful")
	}

	// call the structs
	//fmt.Println("Automigration succesful")
	Db.AutoMigrate(&customers.CustomerDetails{})
	fmt.Println("Automigration succesful")
}
