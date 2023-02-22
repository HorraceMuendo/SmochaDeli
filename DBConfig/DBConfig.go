package database

import (
	_ "SmochaDeliveryApp/Customers/customers"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CustomerConn() *gorm.DB {
	var db *gorm.DB
	DNS := "host= user= password= dbname= port=  sslmode=disabled"

	db, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("connection unsuccesful")
	}
	err = db.AutoMigrate(&Details{})
	if err != nil {
		return
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
