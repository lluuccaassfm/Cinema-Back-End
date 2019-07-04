package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//create struct for room
type Room struct {
	ID			primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Number 		int 				`json:"number" validate:"required,numeric"`
	Type		string 				`json:"type" validate:"required,typeRoom"`
	Capacity	int					`json:"capacity" validate:"required,numeric"`
}

//create struct for update room
type RoomUpdate struct {
	Number 		int 				`json:"number" validate:"required,numeric"`
	Type		string 				`json:"type" validate:"required,typeRoom"`
	Capacity	int					`json:"capacity" validate:"required,numeric"`
}

//create struct for view room in other collection
type RoomView struct {
	Number 		int 				`json:"number" validate:"numeric"`
	Type		string 				`json:"type" validate:"typeRoom"`
	Capacity	int					`json:"capacity" validate:"numeric"`
}




