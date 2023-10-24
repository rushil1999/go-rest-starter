package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-starter/app/controllers"
)


func ImportRoutes(router *gin.Engine) {
	router.GET("/healthCheck", func(c *gin.Context) {
		fmt.Println("Check Sucessfull")
		c.JSON(200, gin.H{"test": "successfull"})
	})


	router.POST("/registerUser", controllers.RegisterUser)
	router.GET ("/users", controllers.GetUsers)
}