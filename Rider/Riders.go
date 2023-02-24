package riders

import "gorm.io/gorm"

type RiderDetails struct {
	gorm.Model

	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Phone       uint   `json:"phone"`
	Location    string `json:"location"`
	Password    string `json:"password"`
	BikeType    string `json:"biketype"`
	PlateNumber string `json:"platenumber"`
}
