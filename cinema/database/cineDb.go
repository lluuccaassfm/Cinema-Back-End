package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"lucas-lima/cinema/models"
)

func InsertCine(r models.Cine) (id interface{}, err error){

	//select cine collection
	c := Db.Collection("cine")

	//insert cine into cine collection
	res, err := c.InsertOne(context.TODO(), r)

	//checks if any error occurs insert an movie
	if err != nil {
		log.Printf("[ERROR] probleming inserting cine: %v", err)
		return
	}

	//assigning value to id
	id = res.InsertedID

	return

}

func DeleteCine(id primitive.ObjectID)(err error){

	//select movie collection
	c := Db.Collection("cine")

	//check if ID exist
	_, err = SearchCineId(id)
	if err != nil {
		return
	}

	//create filter
	filter := bson.D{{"_id",id}}

	//delete cine into cine collection
	_, err = c.DeleteOne(context.TODO(),filter)
	if err != nil {
		log.Printf("[ERROR] Probleming delete cine: %v", err)
		return
	}

	return
}

func SearchCineId(id primitive.ObjectID) (cine models.CineViewInCollection, err error) {
	// select cine collection
	c := Db.Collection("cine")

	// create filter
	filter := bson.D{{"_id", id}}

	//search cine by ID into cine collection
	err = c.FindOne(context.TODO(), filter).Decode(&cine)
	if err != nil {
		log.Printf("[ERROR] probleming searching cine: %v ", err)
		return
	}

	return
}

func SearchAllCine()(cines []models.CineViewInCollection, err error){
	// select cine collection
	c := Db.Collection("cine")

	//search cine into cine collection
	cursor, err := c.Find(context.TODO(),bson.M{})
	if err != nil {
		log.Printf("[ERROR] probleming searching movie: %v ", err)
		return
	}

	//scroll through the list of all cines found
	for cursor.Next(context.TODO()){
		var cine models.CineViewInCollection
		//make cine for json
		_ = cursor.Decode(&cine)
		//add cine to list of cines
		cines = append(cines,cine)
	}

	return
}

func UpdateCine(id primitive.ObjectID, cine models.CineUpdate) (resCine models.CineViewInCollection, err error){

	//select cine collection
	c := Db.Collection("cine")

	//create filter
	filter := bson.M{"_id": id}
	//create update
	update := bson.M{"$set": cine}

	// Update cine into cine collection
	err = c.FindOneAndUpdate(context.TODO(), filter, update).Decode(&resCine)
	if err != nil {
		log.Printf("[ERROR] probleming updating movie: %v\n", err)
		return
	}

	return
}
