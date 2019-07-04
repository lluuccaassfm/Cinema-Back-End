package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"lucas-lima/cinema/models"
)

func InsertReservation(r models.Reservation) (id interface{}, err error){

	//select reservation collection
	c := Db.Collection("reservation")

	//insert reservation into reservation collection
	res, err := c.InsertOne(context.TODO(), r)

	//checks if any error occurs insert an reservation
	if err != nil {
		log.Printf("[ERROR] probleming inserting reservation: %v", err)
		return
	}

	//assigning value to id
	id = res.InsertedID

	return

}

func DeleteReservation(id primitive.ObjectID)(err error){

	//select session collection
	c := Db.Collection("reservation")

	//check if ID exist
	_, err = SearchReservationId(id)
	if err != nil {
		return
	}

	//create filter
	filter := bson.D{{"_id",id}}

	//delete reservation into reservation collection
	_, err = c.DeleteOne(context.TODO(),filter)
	if err != nil {
		log.Printf("[ERROR] Probleming delete reservation: %v", err)
		return
	}

	return
}

func SearchReservationId(id primitive.ObjectID) (reservation models.ReservationViewInCollection, err error) {
	// select reservation collection
	c := Db.Collection("reservation")

	// create filter
	filter := bson.D{{"_id", id}}

	log.Print(id)

	//search reservation by ID into reservation collection
	err = c.FindOne(context.TODO(), filter).Decode(&reservation)
	if err != nil {
		log.Printf("[ERROR] probleming searching reservation: %v ", err)
		return
	}

	return
}

func SearchAllReservation()(reservations []models.ReservationViewInCollection, err error){
	// select reservation collection
	c := Db.Collection("reservation")

	//search reservation into reservation collection
	cursor, err := c.Find(context.TODO(),bson.M{})
	if err != nil {
		log.Printf("[ERROR] probleming searching session: %v ", err)
		return
	}

	//scroll through the list of all reservations found
	for cursor.Next(context.TODO()){
		var reservation models.ReservationViewInCollection
		//make reservation for json
		_ = cursor.Decode(&reservation)
		//add reservation to list of reservations
		reservations = append(reservations,reservation)
	}

	return
}

//func UpdateReservation(id primitive.ObjectID, reservation models.ReservationUpdate) (err error){
//
//	//select reservation collection
//	c := Db.Collection("reservation")
//
//	//create filter
//	filter := bson.M{"_id": id}
//	//create update
//	update := bson.M{"$set": reservation}
//
//	// Update reservation into reservation collection
//	_, err = c.UpdateOne(context.TODO(), filter, update)
//	if err != nil {
//		log.Printf("[ERROR] probleming updating session: %v\n", err)
//		return
//	}
//
//	//message confirmation in log
//	log.Printf("[SUCESS] Reservation Updated ->  %v",reservation)
//	return
//}
