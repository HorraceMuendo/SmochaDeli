package main

// to-do :- get rider by location
// .......  add login func
// .......  add security(Authentication and Authorization)
// .......  fix postgres server

import (
	database "SmochaDeliveryApp/Database"
	routes "SmochaDeliveryApp/Routes"
)

func main() {
	//db connection call
	database.Conn()
	// routes/endponts call
	routes.Routes()

}
