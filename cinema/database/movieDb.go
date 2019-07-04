package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"lucas-lima/cinema/models"
)

func InsertMovie(m models.Movie) (id interface{} ,err error) {

	//select movie collection
	c := Db.Collection("movie")

	//insert movie into movie collection
	res, err := c.InsertOne(context.TODO(), m)

	//checks if any error occurs insert an movie
	if err != nil {
		log.Printf("[ERROR] probleming inserting movie: %v", err)
		return
	}

	//assigning value to id
	id = res.InsertedID

	return
}

func DeleteMovie (id primitive.ObjectID) (err error){

	//select movie collection
	c := Db.Collection("movie")

	//check if ID exist
	_, err = SearchMovieId(id)
	if err != nil {
		return
	}

	//create filter
	filter := bson.D{{"_id",id}}

	//delete movie into movie collection
	_, err = c.DeleteOne(context.TODO(),filter)
	if err != nil {
		log.Printf("[ERROR] Probleming delete movie: %v", err)
		return
	}

	return
}


func SearchMovieId(id primitive.ObjectID)(movie models.Movie, err error){
	// select movie collection
	c := Db.Collection("movie")

	// create filter
	filter := bson.D{{"_id", id}}

	log.Print(id)

	//search movie by ID into movie collection
	err = c.FindOne(context.TODO(), filter).Decode(&movie)

	// checks if any error occurs insert an movie
	if err != nil {
		log.Printf("[ERROR] probleming searching movie: %v ", err)
		return
	}

	return
}

func SearchAllMovie()(movies []models.Movie, err error){
	// select movies collection
	c := Db.Collection("movie")

	// search movies into movie collection
	cursor, err := c.Find(context.TODO(),bson.M{})
	// checks if any error occurs insert an movie
	if err != nil {
		log.Printf("[ERROR] probleming searching movie: %v ", err)
		return
	}

	//scroll through the list of all movies found
	for cursor.Next(context.TODO()){
		var movie models.Movie
		//make movie for json
		_ = cursor.Decode(&movie)
		//add movie to list of movies
		movies = append(movies,movie)
	}

	return
}

func UpdateMovie(id primitive.ObjectID, movie models.MovieUpdate) (resMovie models.Movie,err error){

	// select movies collection
	c := Db.Collection("movie")

	//create filter
	filter := bson.M{"_id": id}
	//create update
	update := bson.M{"$set": movie}

	// Update movie into movie collection
	err = c.FindOneAndUpdate(context.TODO(), filter, update).Decode(&resMovie)
	if err != nil {
		log.Printf("[ERROR] probleming updating movie: %v\n", err)
		return
	}

	return
}