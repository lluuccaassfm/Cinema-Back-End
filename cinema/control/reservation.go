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

// creating function for register a reservation
func CreateReservation(w http.ResponseWriter, r *http.Request){
	//call function for verify request
	reservation, err :=verifyReservation(w,r)
	if err != nil {
		return
	}

	//call function for insert reservation in database
	resId, err := database.InsertReservation(reservation)
	if err != nil {
		configs.ErrorNotFound(w,"reservation")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Reservation Inserted ->  %v",resId)
	//show message success
	_, _ = w.Write([]byte(`{"message":"create success"}`))
}

//creating function for get a reservation
func GetAllReservation(w http.ResponseWriter, r *http.Request){
	//search reservations
	reservations, err := database.SearchAllReservation()
	if err != nil {
		configs.ErrorNotFound(w,"reservations")
		return
	}

	//pass interface{} to []byte
	js, err := json.Marshal(reservations)
	if err != nil {
		log.Printf("[WARN] problem parsing interface to []byte, because: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Reservations found: %v", reservations)
	//show reservations
	_, _ = w.Write(js)
}

//creating function for get a reservation by ID
func GetReservationID(w http.ResponseWriter, r *http.Request){

	//take 'id' into queryString
	query := r.URL.Query().Get("id")
	//create objectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//call function for get reservation in database
	reservation, err := database.SearchReservationId(id)
	if err != nil {
		configs.ErrorNotFound(w,"reservation")
		return
	}

	//pass interface{} to []byte
	js, err := json.Marshal(reservation)
	if err != nil {
		log.Printf("[WARN] problem parsing interface to []byte, because: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Reservation found: %v", reservation)
	//show reservation
	_, _ = w.Write(js)
}

//creating function fo delete a reservation
func DeleteReservation(w http.ResponseWriter, r *http.Request){
	//take 'id' into queryString
	query := r.URL.Query().Get("id")

	//create objectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//delete reservation into reservation collection
	err = database.DeleteReservation(id)
	if err != nil {
		configs.ErrorNotFound(w,"reservation")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Reservation Deleted")
	//show message success
	_, _ = w.Write([]byte(`{"message":"delete success"}`))
}

//creating function for update a reservation
//func UpdateReservation(w http.ResponseWriter, r *http.Request){
//
//	//call function for verify request
//	reservation,err := verifyReservationUpdate(w,r)
//	if err != nil{
//		return
//	}
//
//	//take 'id' into queryString
//	query := r.URL.Query().Get("id")
//	//create objectID
//	id, err :=primitive.ObjectIDFromHex(query)
//
//	//call function for update reservation in database
//	err = database.UpdateReservation(id,reservation)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	//show message success
//	_, _ = w.Write([]byte(`{"message":"update success"}`))
//
//}

//function for verify the request
func verifyReservation( w http.ResponseWriter, r *http.Request)(reservation models.Reservation, err error){

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(r.Body)

	// parses json into reservation struct and checks if any error occurs in json parsing
	if err = json.Unmarshal(bytes,&reservation); err != nil{
		configs.ErrorJson(w,err)
		return
	}

	// validate reservation struct
	if err = validation.Validator.Struct(reservation); err!= nil{
		configs.ErrorValidate(w,err)
		return
	}

	return
}

//function for verify the request
//func verifyReservationUpdate( w http.ResponseWriter, r *http.Request)(reservation models.ReservationUpdate, err error){
//
//	// parsing io.ReadCLoser to slice of bytes []byte
//	bytes, _ := ioutil.ReadAll(r.Body)
//
//	// parses json into reservation struct and checks if any error occurs in json parsing
//	if err = json.Unmarshal(bytes,&reservation); err != nil{
//		log.Printf("[WARN] problem parsing json body, because, %v\n", err)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	// validate reservation struct
//	if err = validation.Validator.Struct(reservation); err!= nil{
//		log.Printf("[WARN] invalid data, because, %v\n", err)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	return
//}