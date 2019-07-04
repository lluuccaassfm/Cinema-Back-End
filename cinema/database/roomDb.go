package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"lucas-lima/cinema/models"
)

func InsertRoom(r models.Room) (id interface{}, err error){

	//select room collection
	c := Db.Collection("room")

	//insert room into room collection
	res, err := c.InsertOne(context.TODO(), r)

	//checks if any error occurs insert an movie
	if err != nil {
		log.Printf("[ERROR] probleming inserting room: %v", err)
		return
	}

	//assigning value to id
	id = res.InsertedID

	return

}

func DeleteRoom(id primitive.ObjectID)(err error){

	//select movie collection
	c := Db.Collection("room")

	//check if ID exist
	_, err = SearchRoomId(id)
	if err != nil {
		return
	}

	//create filter
	filter := bson.D{{"_id",id}}

	//delete room into room collection
	_, err = c.DeleteOne(context.TODO(),filter)
	if err != nil {
		log.Printf("[ERROR] Probleming delete room: %v", err)
		return
	}

	return
}

func SearchRoomId(id primitive.ObjectID) (room models.Room, err error) {
	// select room collection
	c := Db.Collection("room")

	// create filter
	filter := bson.D{{"_id", id}}

	//search room by ID into room collection
	err = c.FindOne(context.TODO(), filter).Decode(&room)
	if err != nil {
		log.Printf("[ERROR] probleming searching room: %v ", err)
		return
	}

	return
}

func SearchAllRoom()(rooms []models.Room, err error){
	// select room collection
	c := Db.Collection("room")

	//search room into room collection
	cursor, err := c.Find(context.TODO(),bson.M{})
	if err != nil {
		log.Printf("[ERROR] probleming searching movie: %v ", err)
		return
	}

	//scroll through the list of all rooms found
	for cursor.Next(context.TODO()){
		var room models.Room
		//make room for json
		_ = cursor.Decode(&room)
		//add room to list of rooms
		rooms = append(rooms,room)
	}

	return
}

func UpdateRoom(id primitive.ObjectID, room models.RoomUpdate) (resRoom models.Room,err error){

	//select room collection
	c := Db.Collection("room")

	//create filter
	filter := bson.M{"_id": id}
	//create update
	update := bson.M{"$set": room}

	// Update room into room collection
	err = c.FindOneAndUpdate(context.TODO(), filter, update).Decode(&resRoom)
	if err != nil {
		log.Printf("[ERROR] probleming updating room: %v\n", err)
		return
	}

	return
}
