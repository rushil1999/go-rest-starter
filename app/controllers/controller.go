package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rest-starter/pkg/models"
)


func SendErrorResponse(context *gin.Context, err error) {
	customError := err.(models.CustomError)
	context.JSON(customError.HttpCode, gin.H{"error": customError.Error()})
}