package control

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"lucas-lima/cinema/configs"
	"lucas-lima/cinema/database"
	"lucas-lima/cinema/models"
	"lucas-lima/cinema/validation"
	"net/http"
)

func CreateRoom(w http.ResponseWriter, r *http.Request){

	//call function for verify request
	room, err :=verifyRoom(w,r)
	if err != nil {
		return
	}

	//call function for insert room in database
	resId, err := database.InsertRoom(room)
	if err != nil {
		configs.ErrorNotFound(w,"room")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Room Inserted ->  %v",resId)

	//show message success
	_, _ = w.Write([]byte(`{"message":"create success"}`))
}


//creating function for get a room
func GetAllRoom(w http.ResponseWriter, r *http.Request){
	//search rooms
	rooms, err := database.SearchAllRoom()
	if err != nil {
		configs.ErrorNotFound(w,"room")
		return
	}

	//pass interface{} to []byte
	js, err := json.Marshal(rooms)
	if err != nil {
		log.Printf("[WARN] problem parsing interface to []byte, because: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Rooms found: %v", rooms)

	//show rooms
	_, _ = w.Write(js)
}

//creating function for get a room by ID
func GetRoomID(w http.ResponseWriter, r *http.Request){

	//take 'id' into queryString
	query := r.URL.Query().Get("id")
	//create objectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//call function for get room in database
	room, err := database.SearchRoomId(id)
	if err != nil {
		configs.ErrorNotFound(w,"room")
		return
	}

	//pass interface{} to []byte
	js, err := json.Marshal(room)
	if err != nil {
		log.Printf("[WARN] problem parsing interface to []byte, because: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Room found: %v", room)

	//show room
	_, _ = w.Write(js)
}

//creating function fo delete a room
func DeleteRoom(w http.ResponseWriter, r *http.Request){
	//take 'id' into queryString
	query := r.URL.Query().Get("id")

	//create objectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//delete room into room collection
	err = database.DeleteRoom(id)
	if err != nil {
		configs.ErrorNotFound(w,"room")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Room Deleted")

	//show message success
	_, _ = w.Write([]byte(`{"message":"delete success"}`))
}

//creating function for update a room
func UpdateRoom(w http.ResponseWriter, r *http.Request){

	//call function for verify request
	room,err := verifyRoomUpdate(w,r)
	if err != nil{
		return
	}

	//take 'id' into queryString
	query := r.URL.Query().Get("id")
	//create objectID
	id, err :=primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//call function for update room in database
	resRoom,err := database.UpdateRoom(id,room)
	if err != nil {
		configs.ErrorNotFound(w,"room")
		return
	}

	//message confirmation in log
	log.Printf("[SUCESS] Room Updated ->  %v",resRoom)

	//show message success
	_, _ = w.Write([]byte(`{"message":"update success"}`))

}
//function for verify the request
func verifyRoom( w http.ResponseWriter, r *http.Request)(room models.Room, err error){

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(r.Body)

	// parses json into room struct and checks if any error occurs in json parsing
	if err = json.Unmarshal(bytes,&room); err != nil{
		configs.ErrorJson(w, err)
		return
	}

	// validate room struct
	if err = validation.Validator.Struct(room); err!= nil{
		configs.ErrorValidate(w,err)
		return
	}

	return
}

//function for verify the request
func verifyRoomUpdate( w http.ResponseWriter, r *http.Request)(room models.RoomUpdate, err error){

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(r.Body)

	// parses json into room struct and checks if any error occurs in json parsing
	if err = json.Unmarshal(bytes,&room); err != nil{
		configs.ErrorJson(w, err)
		return
	}

	// validate room struct
	if err = validation.Validator.Struct(room); err!= nil{
		configs.ErrorValidate(w,err)
		return
	}

	return
}