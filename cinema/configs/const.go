package configs

import (
	"log"
	"net/http"
)

//--------------------
//	Server HTTP
//--------------------
const (

	// defines ip and port address for server instance
	SERVER_ADDR = "172.22.51.134:8080"
	// host for mongo db
	MONGO_HOST = "mongodb://localhost:27017"
)

//-------------------------
//	Paths HTTP
//-------------------------
const (

	USER_PATH 		 = "/user/"
	LOGIN_PATH 		 = "/login/"
	MOVIE_PATH 		 = "/movie/"
	ROOM_PATH 		 = "/room/"
	SESSION_PATH 	 = "/session/"
	CINE_PATH 		 = "/cine/"
	RESERVATION_PATH ="/reservation/"

	)
//-------------------------
// Error Messages
//-------------------------

//shows error when json is invalid
func ErrorJson( w http.ResponseWriter, err error ) {
	log.Printf("[WARN] problem parsing json body, because, %v\n", err)
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(`{"message":"send a valid json"}`))
}

//shows field validation error
func ErrorValidate( w http.ResponseWriter, err error )  {
	log.Printf("[WARN] invalid data, because, %v\n", err)
	w.WriteHeader(http.StatusPreconditionFailed)
	_, _ = w.Write([]byte(`{"message":"some invalid field"}`))
}

//shows error when creating id
func ErrorCreateId( w http.ResponseWriter){
	log.Printf("[ERRO] invalid id information, id must be hex")
	w.WriteHeader(http.StatusPreconditionFailed)
	_, _ = w.Write([]byte(`{"message":"id invalid"}`))
}

//show error when not found in DB
func ErrorNotFound(w http.ResponseWriter, entity string){
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(`{"message":"`+ entity +` not found"}`))
}