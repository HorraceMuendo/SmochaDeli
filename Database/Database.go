package database

import (
	// _ "SmochaDeliveryApp/Customers/customers"
	// "SmochaDeliverApp/customers"
	customers "SmochaDeliveryApp/Customers"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func CustomerConn() *gorm.DB {
	DNS := "host= user= password= dbname= port=  sslmode=disabled"

	db, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("connection unsuccesful")
	}
	// call the structs
	err = db.AutoMigrate(&customers.CustomerDetails{})
	if err != nil {
		log.Fatal("Automigration failed")
	}
	return db
}

//  var db *gorm.DB
// var DNS = "host= user= password= dbname= port=  sslmode=disabled"
// 	// connecting to the database
// 	db, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
// 	if err != nil{
// 		fmt.Println("connection unsuccesful")
// 	}
// 	db.AutoMigrate()
