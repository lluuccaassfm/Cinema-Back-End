package models

type Address struct {
	Logradouro		string	`json:"logradouro" validator:"required"`
	Number			int		`json:"number" validator:"required,numeric"`
	Complement		string	`json:"complement"`
	City			string	`json:"city" validator:"required"`
	State			string	`json:"state" validator:"required"`
}
