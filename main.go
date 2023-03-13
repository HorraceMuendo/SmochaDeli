package main

// to-do :- get rider by location(time function and google maps)
// .......  add login func
// .......  add security(Authentication and Authorization)

import (
	database "SmochaDeliveryApp/Database"
	env "SmochaDeliveryApp/Env"
	routes "SmochaDeliveryApp/Routes"
)

func main() {
	//loading enviroment variables
	env.EnvironmentVar()

	//db connection call
	database.Conn()

	// routes/endponts call
	routes.Routes()

}
