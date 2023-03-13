package database

import (
	"SmochaDeliveryApp/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Conn() {
	var err error

	DNS := os.Getenv("Connstr")

	Db, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("connection unsuccesful")
	}

	Db.AutoMigrate(&model.CustomerDetails{}, &model.RiderDetails{})
	fmt.Println("Automigration succesful")
}
