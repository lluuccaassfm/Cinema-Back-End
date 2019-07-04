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

// creating function for register a cine
func CreateCine(w http.ResponseWriter, r *http.Request){
	//call function for verify request
	cine, err :=verifyCine(w,r)
	if err != nil {
		return
	}

	//call function for insert cine in database
	resId, err := database.InsertCine(cine)
	if err != nil {
		configs.ErrorNotFound(w,"cine")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Cine Inserted ->  %v",resId)
	//show message success
	_, _ = w.Write([]byte(`{"message":"create success"}`))
}

//creating function for get a cine
func GetAllCine(w http.ResponseWriter, r *http.Request){
	//search cines
	cines, err := database.SearchAllCine()
	if err != nil {
		configs.ErrorNotFound(w,"cines")
		return
	}

	//pass interface{} to []byte
	js, err := json.Marshal(cines)
	if err != nil {
		log.Printf("[WARN] problem parsing interface to []byte, because: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Cines found: %v", cines)
	//show cines
	_, _ = w.Write(js)
}

//creating function for get a cine by ID
func GetCineID(w http.ResponseWriter, r *http.Request){

	//take 'id' into queryString
	query := r.URL.Query().Get("id")
	//create objectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//call function for get cine in database
	cine, err := database.SearchCineId(id)
	if err != nil {
		configs.ErrorNotFound(w,"cine")
		return
	}

	//pass interface{} to []byte
	js, err := json.Marshal(cine)
	if err != nil {
		log.Printf("[WARN] problem parsing interface to []byte, because: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Cine found: %v", cine)
	//show cine
	_, _ = w.Write(js)
}

//creating function fo delete a cine
func DeleteCine(w http.ResponseWriter, r *http.Request){
	//take 'id' into queryString
	query := r.URL.Query().Get("id")

	//create objectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//delete cine into cine collection
	err = database.DeleteCine(id)
	if err != nil {
		configs.ErrorNotFound(w,"cine")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Cine Deleted")
	//show message success
	_, _ = w.Write([]byte(`{"message":"delete success"}`))
}

//creating function for update a cine
func UpdateCine(w http.ResponseWriter, r *http.Request){

	//call function for verify request
	cine,err := verifyCineUpdate(w,r)
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

	//call function for update cine in database
	resCine, err := database.UpdateCine(id,cine)
	if err != nil {
		configs.ErrorNotFound(w,"cine")
		return
	}

	//message confirmation in log
	log.Printf("[SUCESS] Cine Updated ->  %v",resCine)
	//show message success
	_, _ = w.Write([]byte(`{"message":"update success"}`))

}
//function for verify the request
func verifyCine( w http.ResponseWriter, r *http.Request)(cine models.Cine, err error){

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(r.Body)

	// parses json into cine struct and checks if any error occurs in json parsing
	if err = json.Unmarshal(bytes,&cine); err != nil{
		configs.ErrorJson(w,err)
		return
	}

	// validate cine struct
	if err = validation.Validator.Struct(cine); err!= nil{
		configs.ErrorValidate(w,err)
		return
	}

	return
}

//function for verify the request
func verifyCineUpdate( w http.ResponseWriter, r *http.Request)(cine models.CineUpdate, err error){

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(r.Body)

	// parses json into cine struct and checks if any error occurs in json parsing
	if err = json.Unmarshal(bytes,&cine); err != nil{
		configs.ErrorJson(w,err)
		return
	}

	// validate cine struct
	if err = validation.Validator.Struct(cine); err!= nil{
		configs.ErrorValidate(w,err)
		return
	}

	return
}