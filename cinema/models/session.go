package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create struct for Session
type Session struct {
	ID		 primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Movie    Movie         		`json:"movie" validate:"required"`
	Room     Room          		`json:"room" validate:"required"`
	Date     string        		`json:"date" validate:"required,date"`
	Duration int				`json:"duration" validate:"required,numeric"`
}

//Create struct for update session
type SessionUpdate struct {
	Date     string  `json:"date" validate:"required,date"`
	Duration int 	 `json:"duration" validate:"required,numeric"`
}

//Create struct for view session in collection
type SessionViewInCollection struct {
	ID       primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie    MovieViewOutCollection `json:"movie"`
	Room     RoomView               `json:"room"`
	Date     string                 `json:"date"`
	Duration int  			        `json:"duration"`
}

//Create struct for View Session in other collection
type SessionViewOutCollection struct {
	Movie    MovieViewOutCollection	`json:"movie"`
	Room     RoomView       		`json:"room"`
	Date     string        			`json:"date"`
	Duration int					`json:"duration"`
}



