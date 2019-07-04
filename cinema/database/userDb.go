package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"lucas-lima/cinema/models"
)

func InsertUser(u models.User) (id interface{} ,err error) {

	//select users collection
	c := Db.Collection("user")

	//insert user into user collection
	res, err := c.InsertOne(context.TODO(), u)

	//checks if any error occurs insert an user
	if err != nil {
		log.Printf("[ERROR] probleming inserting user: %v", err)
		return
	}

	//assigning value to id
	id = res.InsertedID

	return
}

func DeleteUser (id primitive.ObjectID) (res *mongo.DeleteResult,err error){

	//select users collection
	c := Db.Collection("user")

	//check if ID exist
	_, err = SearchUserId(id)
	if err != nil {
		return
	}

	//create filter
	filter := bson.D{{"_id",id}}

	//delete user into user collection
	res, err = c.DeleteOne(context.TODO(),filter)
	if err != nil {
		log.Printf("[ERROR] Probleming delete user: %v", err)
		return
	}

	return
}


func SearchUserId(id primitive.ObjectID)(user models.UserGet, err error){
	// select users collection
	c := Db.Collection("user")

	// create filter
	filter := bson.D{{"_id", id}}

	// try to find user with login
	err = c.FindOne(context.TODO(), filter).Decode(&user)

	// checks if any error occurs insert an user
	if err != nil {
		log.Printf("[ERROR] probleming searching user: %v ", err)
		return
	}

	return
}

func SearchAllUser()(users []models.UserGet, err error){
	// select users collection
	c := Db.Collection("user")

	// try to find user with login
	cursor, err := c.Find(context.TODO(),bson.M{})
	// checks if any error occurs insert an user
	if err != nil {
		log.Printf("[ERROR] probleming searching user: %v ", err)
		return
	}

	//scroll through the list of all users found
	for cursor.Next(context.TODO()){
		var user models.UserGet
		//make user for json
		_ = cursor.Decode(&user)
		//add user to list of users
		users = append(users,user)
	}

	return
}

func SearchLoginUser(login models.LoginUser)(user models.User,err error){
	// select users collection
	c := Db.Collection("user")

	// create filter
	filter := bson.D{{"email", login.Email}, {"pass", login.Pass}}

	// try to find user with login
	err = c.FindOne(context.TODO(), filter).Decode(&user)
	// checks if any error occurs insert an user
	if err != nil {
		log.Printf("[ERROR] probleming searching user: %v", err)
		return
	}

	return
}

func UpdateUser(id primitive.ObjectID, user models.UserUpdate) (resUser models.User, err error){

	// select users collection
	c := Db.Collection("user")

	//create filter
	filter := bson.M{"_id": id}
	//create update
	update := bson.M{"$set": user}

	// Update user into user collection
	//resUser, err = c.UpdateOne(context.TODO(), filter, update)

	// Find and Update user into user collection
	err = c.FindOneAndUpdate(context.TODO(), filter, update).Decode(&resUser)
	if err != nil {
		log.Printf("[ERROR] probleming updating user: %v\n", err)
		return
	}

	return
}

func UpdateUserSpecific(id primitive.ObjectID, user models.UserUpdateSpecific) (resUser models.User,err error){

	// select users collection
	c := Db.Collection("user")

	//create filter
	filter := bson.M{"_id": id}
	//create update
	update := bson.M{"$set": user}

	// Update user into user collection
	err = c.FindOneAndUpdate(context.TODO(), filter, update).Decode(&resUser)
	if err != nil {
		log.Printf("[ERROR] probleming updating user: %v\n", err)
		return
	}

	return
}

