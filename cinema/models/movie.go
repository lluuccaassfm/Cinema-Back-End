package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//create struct for movie
type Movie struct{
	ID 		  		primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Name			string				`json:"name" validate:"required"`
	Genre			string				`json:"genre" validate:"required,genre"`
	Classification 	int					`json:"classification" validate:"numeric"`
	Duration		int					`json:"duration" validate:"required,numeric"`
	Description 	string				`json:"description" validate:"required"`
	Situation		string				`json:"situation" validate:"required,situation"`
	Type			string				`json:"type" validate:"required,typeMovie"`
}

//create struct for movie update
type MovieUpdate struct{
	Name			string				`json:"name" validate:"required"`
	Genre			string				`json:"genre" validate:"required,genre"`
	Classification 	int					`json:"classification" validate:"numeric"`
	Duration		int					`json:"duration" validate:"required,numeric"`
	Description 	string				`json:"description" validate:"required"`
	Situation		string				`json:"situation" validate:"required,situation"`
	Type			string				`json:"type" validate:"required,typeMovie"`
}

//create struct for view movie in other collection
type MovieViewOutCollection struct{
	Name			string				`json:"name"`
	Genre			string				`json:"genre"`
	Classification 	int					`json:"classification"`
	Duration		int					`json:"duration"`
	Description 	string				`json:"description"`
	Situation		string				`json:"situation"`
	Type			string				`json:"type"`
}


