package main

// to-do :- get rider by location(time function and google maps)
// .......  add login func
// .......  add security(Authentication and Authorization)

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
