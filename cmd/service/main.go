package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-starter/app/api"
	"go-rest-starter/app/middlewares"
)


func main() {
	router := gin.Default()
	router.Use(middlewares.LogAPIRequest())
	api.ImportRoutes(router)
	fmt.Println("Server Started")
	router.Run("localhost:8000")

}