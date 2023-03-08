package database

import (
	"SmochaDeliveryApp/model"
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

	Db.AutoMigrate(&model.CustomerDetails{}, &model.RiderDetails{})
	fmt.Println("Automigration succesful")
}
