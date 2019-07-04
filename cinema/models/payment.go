package models

type Payment struct {
	Name  	string	`json:"name" validator:"required"`
	Type  	string	`json:"type" validator:"required, typePayment"`
	Value	float32	`json:"value" validator:"required"`
	Date 	string	`json:"date" validator:"required,date"`
}