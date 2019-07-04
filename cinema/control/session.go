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

// creating function for register a session
func CreateSession(w http.ResponseWriter, r *http.Request){
	//call function for verify request
	session, err :=verifySession(w,r)
	if err != nil {
		return
	}

	//call function for insert session in database
	resId, err := database.InsertSession(session)
	if err != nil {
		configs.ErrorNotFound(w,"session")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Session Inserted ->  %v",resId)
	//show message success
	_, _ = w.Write([]byte(`{"message":"create success"}`))
}

//creating function for get a session
func GetAllSession(w http.ResponseWriter, r *http.Request){
	//search sessions
	sessions, err := database.SearchAllSession()
	if err != nil {
		configs.ErrorNotFound(w,"sessions")
		return
	}

	//pass interface{} to []byte
	js, err := json.Marshal(sessions)
	if err != nil {
		log.Printf("[WARN] problem parsing interface to []byte, because: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Sessions found: %v", sessions)
	//show sessions
	_, _ = w.Write(js)
}

//creating function for get a session by ID
func GetSessionID(w http.ResponseWriter, r *http.Request){

	//take 'id' into queryString
	query := r.URL.Query().Get("id")
	//create objectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//call function for get session in database
	session, err := database.SearchSessionId(id)
	if err != nil {
		configs.ErrorNotFound(w,"session")
		return
	}

	//pass interface{} to []byte
	js, err := json.Marshal(session)
	if err != nil {
		log.Printf("[WARN] problem parsing interface to []byte, because: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Session found: %v", session)

	//show session
	_, _ = w.Write(js)
}

//creating function fo delete a session
func DeleteSession(w http.ResponseWriter, r *http.Request){
	//take 'id' into queryString
	query := r.URL.Query().Get("id")

	//create objectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//delete session into session collection
	err = database.DeleteSession(id)
	if err != nil {
		configs.ErrorNotFound(w,"session")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Session Deleted")
	//show message success
	_, _ = w.Write([]byte(`{"message":"delete success"}`))
}

//creating function for update a session
func UpdateSession(w http.ResponseWriter, r *http.Request){

	//call function for verify request
	session,err := verifySessionUpdate(w,r)
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

	//call function for update session in database
	resSession, err := database.UpdateSession(id,session)
	if err != nil {
		configs.ErrorNotFound(w,"session")
		return
	}

	//message confirmation in log
	log.Printf("[SUCESS] Session Updated ->  %v",resSession)
	//show message success
	_, _ = w.Write([]byte(`{"message":"update success"}`))

}
//function for verify the request
func verifySession( w http.ResponseWriter, r *http.Request)(session models.Session, err error){

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(r.Body)

	// parses json into session struct and checks if any error occurs in json parsing
	if err = json.Unmarshal(bytes,&session); err != nil{
		configs.ErrorJson(w,err)
		return
	}

	// validate session struct
	if err = validation.Validator.Struct(session); err!= nil{
		configs.ErrorValidate(w,err)
		return
	}

	return
}

//function for verify the request
func verifySessionUpdate( w http.ResponseWriter, r *http.Request)(session models.SessionUpdate, err error){

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(r.Body)

	// parses json into session struct and checks if any error occurs in json parsing
	if err = json.Unmarshal(bytes,&session); err != nil{
		configs.ErrorJson(w,err)
		return
	}

	// validate session struct
	if err = validation.Validator.Struct(session); err!= nil{
		configs.ErrorValidate(w,err)
		return
	}

	return
}