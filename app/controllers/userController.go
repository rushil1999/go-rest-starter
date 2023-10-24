package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rest-starter/pkg/models"
	"go-rest-starter/app/services"
)

func RegisterUser(context *gin.Context) {
	var person models.Person
	if err := context.ShouldBindJSON(&person); err != nil {
		customError := models.CustomError{
			Message: "Internal Server Error", 
			DebugMessage: "could not bind json", 
			HttpCode: 500,
		}
		SendErrorResponse(context, customError)
		return
	}
	newPerson, err := services.RegisterUser(person)
	if err != nil {
		SendErrorResponse(context, err)
		return 
	}
	context.JSON(200, newPerson)
}

func GetUsers(context *gin.Context) {
	context.JSON(200, models.UserData)
}