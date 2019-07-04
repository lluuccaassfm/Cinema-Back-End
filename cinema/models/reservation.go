package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create struct for reservation
type Reservation struct {
	ID		primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	User 	UserReservation		`json:"user" validate:"required"`
	Cine	CineReservation		`json:"cine" validate:"required"`
	Session Session				`json:"session" validate:"required"`
	Payment	Payment				`json:"payment" validate:"required"`
	Seat	string				`json:"seat" validate:"required"`
}

//Create struct for view reservation in collection
type ReservationViewInCollection  struct {
	ID		primitive.ObjectID			`json:"_id,omitempty" bson:"_id,omitempty"`
	User 	UserReservation				`json:"user"`
	Cine	CineReservation				`json:"cine"`
	Session SessionViewOutCollection	`json:"session"`
	Payment	Payment						`json:"payment"`
	Seat	string						`json:"seat"`
}