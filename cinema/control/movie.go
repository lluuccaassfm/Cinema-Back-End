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

// creating function for register a movie
func CreateMovie(w http.ResponseWriter, r *http.Request){

	//call function for verify request
	movie,err := verifyMovie(w,r)
	if err != nil{
		return
	}

	//call function for insert movie in database
	id, err := database.InsertMovie(movie)
	if err != nil {
		configs.ErrorNotFound(w,"movie")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Movie Inserted ->  %v",id)
	//show message success
	_, _ = w.Write([]byte(`{"message":"create success"}`))
}

//creating function for get a movie
func GetAllMovie(w http.ResponseWriter, r *http.Request){
	//search movies
	movies, err := database.SearchAllMovie()
	if err != nil {
		configs.ErrorNotFound(w,"movies")
		return
	}

	//pass interface{} to []byte
	js, err := json.Marshal(movies)
	if err != nil {
		log.Printf("[WARN] problem parsing interface to []byte, because: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Movies found: %v", movies)

	//show movies
	_, _ = w.Write(js)
}

//creating function for get a movie by ID
func GetMovieID(w http.ResponseWriter, r *http.Request){

	//take 'id' into queryString
	query := r.URL.Query().Get("id")
	//create objectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//call function for get movie in database
	movie, err := database.SearchMovieId(id)
	if err != nil {
		configs.ErrorNotFound(w,"movie")
		return
	}

	//pass interface{} to []byte
	js, err := json.Marshal(movie)
	if err != nil {
		log.Printf("[WARN] problem parsing interface to []byte, because: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Movie found: %v", movie)

	//show movie
	_, _ = w.Write(js)
}

//creating function fo delete a movie
func DeleteMovie(w http.ResponseWriter, r *http.Request){
	//take 'id' into queryString
	query := r.URL.Query().Get("id")

	//create objectID
	id,err := primitive.ObjectIDFromHex(query)
	if err != nil {
		configs.ErrorCreateId(w)
		return
	}

	//delete movie into movie collection
	err = database.DeleteMovie(id)
	if err != nil {
		configs.ErrorNotFound(w,"movie")
		return
	}

	//message confirmation in log
	log.Printf("[SUCCESS] Movie Deleted")

	//show message success
	_, _ = w.Write([]byte(`{"message":"delete success"}`))
}

//creating function for update a movie
func UpdateMovie(w http.ResponseWriter, r *http.Request){

	//call function for verify request
	movie,err := verifyMovieUpdate(w,r)
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

	//call function for update movie in database
	resMovie,err := database.UpdateMovie(id,movie)
	if err != nil {
		configs.ErrorNotFound(w,"movie")
		return
	}

	//message confirmation in log
	log.Printf("[SUCESS] Movie Updated ->  %v",resMovie)

	//show message success
	_, _ = w.Write([]byte(`{"message":"update success"}`))

}

//function for verify is body
func verifyMovie(w http.ResponseWriter, r *http.Request)(movie models.Movie, err error){

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(r.Body)

	// parses json into movie struct and checks if any error occurs in json parsing
	if err = json.Unmarshal(bytes,&movie); err != nil{
		configs.ErrorJson(w,err)
		return
	}

	// validate movie struct
	if err = validation.Validator.Struct(movie); err!= nil{
		configs.ErrorValidate(w,err)
		return
	}
	return
}

//function for verify is body
func verifyMovieUpdate(w http.ResponseWriter, r *http.Request)(movie models.MovieUpdate, err error){

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(r.Body)

	// parses json into movie struct and checks if any error occurs in json parsing
	if err = json.Unmarshal(bytes,&movie); err != nil{
		configs.ErrorJson(w,err)
		return
	}

	// validate movie struct
	if err = validation.Validator.Struct(movie); err!= nil{
		configs.ErrorValidate(w,err)
		return
	}
	return
}
