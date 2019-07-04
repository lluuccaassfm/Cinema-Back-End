package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//create struct for user
type User struct {
	ID 		  primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string  				`json:"email" validate:"required,email"`
	Pass      string  				`json:"pass"  validate:"required,max=10,min=8,alphanum"`
	Name      string  				`json:"name"  validate:"required,min=3,max=40"`
	Sex       string  				`json:"sex"   validate:"required,sex"`
	BirthDate string  				`json:"birthDate" validate:"required,date"`
	Address   Address 				`json:"address" validate:"required"`
}

//create struct for get user
type UserGet struct {
	ID 		  primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string  				`json:"email"`
	Name      string  				`json:"name"`
	Sex       string  				`json:"sex"`
	BirthDate string  				`json:"birthDate"`
	Address   Address 				`json:"address"`
}

//create struct for login of user
type LoginUser struct {
	Email   string        `json:"email" validate:"required,email"`
	Pass    string        `json:"pass" validate:"required,max=12,min=8,alphanum"`
}

//create struct for update user
type UserUpdate struct {
	Name      string  	`json:"name"  validate:"required,min=3,max=40"`
	Sex       string  	`json:"sex"   validate:"required,sex"`
	BirthDate string  	`json:"birthDate" validate:"required,date"`
	Address   Address 	`json:"address" validate:"required"`
}

//create struct for update email and pass user
type UserUpdateSpecific struct {
	Email   string  `json:"email" validate:"omitempty,email"`
	Pass	string  `json:"pass"  validate:"omitempty,max=10,min=8,alphanum"`
}

//create struct for reservation of user
type UserReservation struct {
	Email     string  	`json:"email" validate:"required,email"`
	Name      string  	`json:"name"  validate:"required,min=3,max=40"`
	BirthDate string  	`json:"birthDate" validate:"required,date"`
}

