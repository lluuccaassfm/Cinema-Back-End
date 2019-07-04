package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"lucas-lima/cinema/models"
)

func InsertSession(r models.Session) (id interface{}, err error){

	//select session collection
	c := Db.Collection("session")

	//insert session into session collection
	res, err := c.InsertOne(context.TODO(), r)

	//checks if any error occurs insert an session
	if err != nil {
		log.Printf("[ERROR] probleming inserting session: %v", err)
		return
	}

	//assigning value to id
	id = res.InsertedID

	return

}

func DeleteSession(id primitive.ObjectID)(err error){

	//select session collection
	c := Db.Collection("session")

	//check if ID exist
	_, err = SearchSessionId(id)
	if err != nil {
		return
	}

	//create filter
	filter := bson.D{{"_id",id}}

	//delete session into session collection
	_, err = c.DeleteOne(context.TODO(),filter)
	if err != nil {
		log.Printf("[ERROR] Probleming delete session: %v", err)
		return
	}

	return
}

func SearchSessionId(id primitive.ObjectID) (session models.SessionViewInCollection, err error) {
	// select session collection
	c := Db.Collection("session")

	// create filter
	filter := bson.D{{"_id", id}}

	log.Print(id)

	//search session by ID into session collection
	err = c.FindOne(context.TODO(), filter).Decode(&session)
	if err != nil {
		log.Printf("[ERROR] probleming searching session: %v ", err)
		return
	}

	return
}

func SearchAllSession()(sessions []models.SessionViewInCollection, err error){
	// select session collection
	c := Db.Collection("session")

	//search session into session collection
	cursor, err := c.Find(context.TODO(),bson.M{})
	if err != nil {
		log.Printf("[ERROR] probleming searching session: %v ", err)
		return
	}

	//scroll through the list of all sessions found
	for cursor.Next(context.TODO()){
		var session models.SessionViewInCollection
		//make session for json
		_ = cursor.Decode(&session)
		//add session to list of sessions
		sessions = append(sessions,session)
	}

	return
}

func UpdateSession(id primitive.ObjectID, session models.SessionUpdate) (resSession models.SessionViewInCollection, err error){

	//select session collection
	c := Db.Collection("session")

	//create filter
	filter := bson.M{"_id": id}
	//create update
	update := bson.M{"$set": session}

	// Update session into session collection
	err = c.FindOneAndUpdate(context.TODO(), filter, update).Decode(&resSession)
	if err != nil {
		log.Printf("[ERROR] probleming updating session: %v\n", err)
		return
	}

	return
}
