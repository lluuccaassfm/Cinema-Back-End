package main

import (
	"lucas-lima/cinema/database"
	"lucas-lima/cinema/service"
)

func main(){

	// assure server closing at end of execution
	defer service.StopServer()

	// call db client constructor
	database.CreateClient()

	// call start server function
	service.StartSever()
}


