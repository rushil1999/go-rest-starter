package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
	"log"
)

func LogAPIRequest() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		log.Println("Request received at time stamp:", t)
		log.Println(context.ClientIP())
		log.Println(context.ContentType())
		context.Next()
		latency := time.Since(t)
		log.Println("Time to execute:", latency)
	}
}