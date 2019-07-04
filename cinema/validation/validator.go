package validation

import (
	val "gopkg.in/go-playground/validator.v9"
	"regexp"
)

// create a global validator
var Validator *val.Validate


// function to instanceate a validator
func CreateValidator() {

	//call create validator
	Validator = val.New()

	//create new validator for sex
	_ = Validator.RegisterValidation("sex", ValidatorSex)

	//create new validator for situation of movie
	_ = Validator.RegisterValidation("situation", ValidatorSituationMovie)

	//create new validator for types of movies
	_ = Validator.RegisterValidation("genre",ValidatorGenreMovie)

	//create new validator for date
	_ = Validator.RegisterValidation("date", ValidatorDate)

	//create new validator for type of room
	_ = Validator.RegisterValidation("typeRoom", ValidatorTypeRoom)
	
	//create new validator for type of payment	
	_= Validator.RegisterValidation("typePayment", ValidatorTypePayment)

	//create new validator for type of movie
	_= Validator.RegisterValidation("typeMovie", ValidatorTypeMovie)
}

//function to validate sex
func ValidatorSex(field val.FieldLevel) bool{

	//creates a string of the passed value
	sex := field.Field().String()
	//verify the past value
	return sex == "male" || sex == "female"
}

// function to validate situation of movie
func ValidatorSituationMovie(field val.FieldLevel) bool {
	//creates a string of the passed value
	situation := field.Field().String()
	//verify the past value
	return  situation == "em cartaz" || situation == "em breve"
}

//function to validate type of movie
func ValidatorGenreMovie(field val.FieldLevel) bool {

	//creates a string of the passed value
	genre := field.Field().String()
	//verify the past value
	return  genre == "action" || genre == "comedy" || genre == "drama" ||
			genre == "science fiction" || genre == "animation" || genre == "adventure" ||
			genre == "documentary" || genre == "war" || genre == "musical" || genre == "thriller" ||
			genre == "suspense" || genre == "terror"

}

//function to validate for date
func ValidatorDate(field val.FieldLevel) bool{

	//creates a string of the passed value
	date := field.Field().String()
	//create regex of date
	re := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])")
	//verify the past value
	return re.MatchString(date)
}

//function to validate situation
func ValidatorTypeRoom(field val.FieldLevel) bool {

	//creates a string of the passed value
	typeRoom := field.Field().String()
	//verify the past value
	return typeRoom == "2D"	|| typeRoom == "3D" || typeRoom == "3D/XD"
}

//function to validate type
func ValidatorTypePayment(field val.FieldLevel) bool {

	//creates a string of the passed value
	typePayment := field.Field().String()
	//verify the past value
	return typePayment == "card" || typePayment == "banking billet " || typePayment == "PicPay"
}

func ValidatorTypeMovie(field val.FieldLevel) bool {

	//creates a string of the passed value
	typeMovie := field.Field().String()
	//verify the past value
	return typeMovie == "subtitled" || typeMovie == "dubbed"
}
