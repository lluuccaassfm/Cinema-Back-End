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

// function for register a user
func CreateUser(w http.ResponseWriter, r *http.Request){
	//call function for verify and validation of request
	user, err := verifyUser(w,r)

	if err != nil{
		return
	}

	//call function for insert user in database
	id, err := database.InsertUser(user)
	if err != nil {
		configs.ErrorNotFound(w,"user")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] User Inserted ->  %v",id)
	//show message success
	_, _ = w.Write([]byte(`{"message":"create success"}`))
}

// function for delete a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//take 'id' into queryString
	query := r.URL.Query().Get("id")

	//create objectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//delete user into user collection
	_,err = database.DeleteUser(id)
	if err != nil {
		configs.ErrorNotFound(w,"user")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] User Deleted")

	//show message success
	_, _ = w.Write([]byte(`{"message":"delete success"}`))
}

//function for get a user for ID
func GetUserId(w http.ResponseWriter, r *http.Request) {

	//take 'id' into queryString
	query := r.URL.Query().Get("id")

	//create objectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//search user by ID
	user, err := database.SearchUserId(id)
	if err != nil {
		configs.ErrorNotFound(w,"user")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] User found: %v", user)

	//pass interface{} to []byte
	js, err := json.Marshal(user)
	if err != nil {
		log.Printf("[WARN] problem parsing interface to []byte, because: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//show user
	_, _ = w.Write(js)
}

//function for get all users
func GetAllUser(w http.ResponseWriter, r *http.Request) {

	//search users
	users, err := database.SearchAllUser()
	if err != nil {
		configs.ErrorNotFound(w,"users")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Users found: %v", users)

	//pass interface{} to []byte
	js, err := json.Marshal(users)
	if err != nil {
		log.Printf("[WARN] problem parsing interface to []byte, because: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//show users
	_, _ = w.Write(js)
}


// function for update a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	//call function for verify and validation of request
	userUpdate, err := verifyUserUpdate(w,r)
	if err != nil{
		return
	}

	//take 'id' into queryString
	query := r.URL.Query().Get("id")

	//create ObjectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//call function for insert user in database
	resUser, err := database.UpdateUser(id,userUpdate)
	if err != nil{
		configs.ErrorNotFound(w,"user")
		return
	}

	//message confirmation in log
	log.Printf("[SUCESS] User Updated ->  %v",resUser)
	//show message sucess
	_, _ = w.Write([]byte(`{"message":" update success"}`))
}

// function for update pass and email of user
func UpdateUserSpecific(w http.ResponseWriter, r *http.Request){

	//call function for verify and validation of request
	userUpdate, err := verifyUserUpdateSpecific(w,r)
	if err != nil{
		return
	}

	//take 'id' into queryString
	query := r.URL.Query().Get("id")

	//create id ObjectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//call function for insert user in database
	user,err := database.UpdateUserSpecific(id,userUpdate)
	if err != nil {
		configs.ErrorNotFound(w,"user")
		return
	}

	//message confirmation in log
	log.Printf("[SUCESS] User Updated ->  %v",user)

	//show message success
	_, _ = w.Write([]byte(`{"message":"update success"}`))
}

//function for verify data of login the user
func LoginUser(w http.ResponseWriter, r *http.Request){

	// declare use entity
	var login models.LoginUser

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(r.Body)

	// parses json into user struct
	err := json.Unmarshal(bytes, &login)

	// checks if any error occurs in json parsing
	if err != nil {
		configs.ErrorJson(w,err)
		return
	}

	// checks if struct is a valid one
	if err := validation.Validator.Struct(login); err != nil {
		configs.ErrorValidate(w,err)
		return
	}

	//verify if login exist
	user, err := database.SearchLoginUser(login)
	if err != nil {
		configs.ErrorNotFound(w,"user")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] User found: %v", user)

	//show message success for user
	_, _ = w.Write([]byte(`{"message":"login success"}`))
}


func verifyUser( w http.ResponseWriter, r *http.Request)(user models.User, err error){

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(r.Body)

	// parses json into user struct
	if err = json.Unmarshal(bytes, &user); err != nil {
		configs.ErrorJson(w,err)
		return
	}

	// verify is validator by user struct
	if err = validation.Validator.Struct(user); err != nil{
		configs.ErrorValidate(w,err)
		return
	}

	return
}

func verifyUserUpdate(w http.ResponseWriter, r *http.Request)(userUpdate models.UserUpdate ,err error){

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(r.Body)

	// parses json into user struct
	if err = json.Unmarshal(bytes, &userUpdate); err != nil {
		configs.ErrorJson(w,err)
		return
	}

	// verify is validator by user struct
	if err = validation.Validator.Struct(userUpdate); err != nil{
		configs.ErrorValidate(w,err)
		return
	}

	return
}

func verifyUserUpdateSpecific(w http.ResponseWriter, r *http.Request)(userUpdate models.UserUpdateSpecific ,err error){

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(r.Body)

	// parses json into user struct
	if err = json.Unmarshal(bytes, &userUpdate); err != nil {
		configs.ErrorJson(w,err)
		return
	}

	// verify is validator by user struct
	if err = validation.Validator.Struct(userUpdate); err != nil{
		configs.ErrorValidate(w,err)
		return
	}

	return
}