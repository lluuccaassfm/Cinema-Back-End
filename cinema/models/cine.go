package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//create struct for cine
type Cine struct {
	ID 		 primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string    			`json:"name" validate:"required"`
	Address  Address   			`json:"address" validate:"required"`
	Rooms    []Room          	`json:"rooms" validate:"required"`
	Movies   []Movie    		`json:"movies" validate:"required"`
	Sessions []Session 			`json:"sessions" validate:"required"`
}

//create struct for update cine
type CineUpdate struct {
	Name     string    	`json:"name" validate:"required"`
	Address  Address   	`json:"address" validate:"required"`
	Rooms    []Room     `json:"rooms" validate:"required"`
	Sessions []Session 	`json:"sessions" validate:"required"`
	Movies   []Movie    `json:"movies" validate:"required"`
}

//Create struct for View Cine in collection
type CineViewInCollection struct {
	ID 		 primitive.ObjectID		    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string    				    `json:"name"`
	Address  Address   				    `json:"address"`
	Rooms    []RoomView          	    `json:"rooms"`
	Sessions []SessionViewOutCollection `json:"sessions"`
	Movies   []MovieViewOutCollection   `json:"movies"`
}

//Create struct for entity 'reservation'
type CineReservation struct {
	Name    string    `json:"name" validator:"required"`
	Address Address   `json:"address" validator:"required"`
}

