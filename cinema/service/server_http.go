package service

import (
	"log"
	"lucas-lima/cinema/configs"
	"lucas-lima/cinema/validation"
	"net/http"
	"time"
)

func StartSever(){

	// creates a handler/router
	h	:= createHandler()

	// creates a HTTP server with default parameters
	server := createServer()

	// creates a CORS struct to deal with cross origin
	c := createCORS()

	// set cors treatment on router
	handler := c.Handler(h)

	// associate handler to server
	server.Handler = handler

	// create global validator
	validation.CreateValidator()

	// instanciates a HTTP server wrapped in a log fatal
	log.Fatal(server.ListenAndServe())

}

func StopServer(){}

func createServer () (server *http.Server){

	server = &http.Server{

		Addr:         configs.SERVER_ADDR,
		IdleTimeout:  100 * time.Millisecond,
		ReadTimeout:  100 * time.Millisecond,
		WriteTimeout: 100 * time.Millisecond,
	}

	return
}
