package services

import (
	"fmt"
	"github.com/google/uuid"
	"go-rest-starter/pkg/models"
)



func RegisterUser(person models.Person) (models.Person, error ){ 
	err := validateNewUser(person)
	if err != nil {
		return  models.Person{}, err
	}
	newUserId := uuid.New().String()
	person.Id = newUserId
	externalApiResponse := GetUniqueBankAccountNumber()
	creditCardDetails := externalApiResponse["credit_card"].(map[string]interface{})
	accountNumber := creditCardDetails["cc_number"]
	person.UserAccount = accountNumber.(string)
	models.UserData = append(models.UserData, person)
	return person, nil
}


func GetUsers() []models.Person { 
	return models.UserData
}

func validateNewUser(person models.Person) error {
	isStringEmpty := CheckEmptyString(person.Name)
	if isStringEmpty {
		return models.CustomError{
			Message: "Invalid Person Name", DebugMessage: "Person name cannot be of length 0",
		}
	}
	isStringEmpty = CheckEmptyString(person.Password)
	if isStringEmpty {
		return models.CustomError{
			Message: "Invalid Password ", DebugMessage: "Person password cannot be of length 0",
		}
	}

	isStringEmpty = CheckEmptyString(person.Email)
	if isStringEmpty {
		return models.CustomError{
			Message: "Invalid Email ", DebugMessage: "Person email cannot be of length 0",
		}
	}
	return nil
}