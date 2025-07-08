package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello Stranger!"})
}
